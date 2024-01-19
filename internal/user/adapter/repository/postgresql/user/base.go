package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/google/wire"

	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"

	userPgQeries "github.com/duongnln96/blog-realworld/internal/user/infras/postgresql/user"
	psqlAdapter "github.com/duongnln96/blog-realworld/pkg/adapter/postgres"
)

type repoManager struct {
	querier *userPgQeries.Queries
	pg      psqlAdapter.PostgresDBAdapterI
}

var _ port.UserRepoI = (*repoManager)(nil)

var RepositorySet = wire.NewSet(NewRepoManager)

func NewRepoManager(pg psqlAdapter.PostgresDBAdapterI) port.UserRepoI {
	querier := userPgQeries.New(pg.GetDB())
	return &repoManager{
		querier: querier,
		pg:      pg,
	}
}

func (r *repoManager) Create(ctx context.Context, req domain.User) (domain.User, error) {

	dbUser, err := r.querier.CreateUser(ctx, userPgQeries.CreateUserParams{
		Username:     req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
		Status:       req.Status.ToString(),
	})
	if err != nil {
		return domain.User{}, fmt.Errorf("querier.CreateUser %s", err.Error())
	}

	return domain.User{
		ID:       dbUser.ID,
		Name:     dbUser.Username,
		Email:    dbUser.Email,
		Password: dbUser.PasswordHash,
		Bio:      dbUser.Bio,
		Status:   domain.NewUserStatusFromString(dbUser.Status),

		CreatedDate: &dbUser.CreatedDate,
		UpdatedDate: &dbUser.UpdatedDate.Time,
	}, nil
}

func (r *repoManager) Update(ctx context.Context, req domain.User) (domain.User, error) {

	tx, err := r.pg.GetDB().Begin()
	if err != nil {
		return domain.User{}, fmt.Errorf("pg.GetDB().Begin() %s", err.Error())
	}

	qtx := r.querier.WithTx(tx)

	result, err := qtx.UpdateUser(ctx, userPgQeries.UpdateUserParams{
		ID:           req.ID,
		Email:        req.Email,
		Username:     req.Name,
		PasswordHash: req.Password,
		Bio:          req.Bio,
		Status:       req.Status.ToString(),
	})

	return domain.User{
		ID:       result.ID,
		Name:     result.Username,
		Email:    result.Email,
		Password: result.PasswordHash,
		Bio:      result.Bio,
		Status:   domain.NewUserStatusFromString(result.Status),

		CreatedDate: &result.CreatedDate,
		UpdatedDate: &result.UpdatedDate.Time,
	}, tx.Commit()
}

func (r *repoManager) GetOneByID(ctx context.Context, userID uuid.UUID) (domain.User, error) {

	rowResult, err := r.querier.GetOneByID(ctx, userID)
	if err != nil && err != sql.ErrNoRows {
		return domain.User{}, fmt.Errorf("querier.GetOneByID %s", err.Error())
	}

	return domain.User{
		ID:     rowResult.ID,
		Name:   rowResult.Username,
		Bio:    rowResult.Bio,
		Status: domain.NewUserStatusFromString(rowResult.Status),

		CreatedDate: &rowResult.CreatedDate,
		UpdatedDate: &rowResult.UpdatedDate.Time,
	}, nil
}

func (r *repoManager) GetOneByEmail(ctx context.Context, email string) (user domain.User, err error) {

	rowResult, err := r.querier.GetOneByEmail(ctx, email)
	if err != nil && err != sql.ErrNoRows {
		return domain.User{}, fmt.Errorf("querier.GetOneByEmail %s", err.Error())
	}

	return domain.User{
		ID:     rowResult.ID,
		Name:   rowResult.Username,
		Bio:    rowResult.Bio,
		Status: domain.NewUserStatusFromString(rowResult.Status),
	}, nil
}
