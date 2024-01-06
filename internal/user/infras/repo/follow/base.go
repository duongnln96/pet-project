package follow

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/google/wire"

	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"

	followPgQeries "github.com/duongnln96/blog-realworld/internal/user/infras/postgresql/follow"
	psqlAdapter "github.com/duongnln96/blog-realworld/pkg/adapter/postgres"
)

const (
	MAX_DEFAULT_LIMIT_REPO int32 = 100
)

type repoManager struct {
	querier *followPgQeries.Queries
	pg      psqlAdapter.PostgresDBAdapterI
}

var _ port.FollowRepoI = (*repoManager)(nil)

var RepositorySet = wire.NewSet(NewRepoManager)

func NewRepoManager(pg psqlAdapter.PostgresDBAdapterI) port.FollowRepoI {
	querier := followPgQeries.New(pg.GetDB())
	return &repoManager{
		querier: querier,
		pg:      pg,
	}
}

func (r *repoManager) Create(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (domain.Follow, error) {
	row, err := r.querier.CreateFollow(ctx, followPgQeries.CreateFollowParams{
		FollowedUserID:  followedUserID,
		FollowingUserID: followingUserID,
	})
	if err != nil {
		return domain.Follow{}, errors.New(fmt.Sprintf("querier.CreateFollow %s", err.Error()))
	}

	return domain.Follow{
		ID:              row.ID,
		FollowedUserID:  row.FollowedUserID,
		FollowingUserID: row.FollowingUserID,
		CreatedDate:     &row.CreatedDate,
		UpdatedDate:     &row.UpdatedDate.Time,
	}, nil
}

func (r *repoManager) AllByFollowedUserID(ctx context.Context, followedUserID uuid.UUID, offset int64, limit int32) (domain.Follows, error) {

	queryReq := followPgQeries.AllByFollowedUserIDParams{
		FollowedUserID: followedUserID,
		Column2:        []string{domain.ActiveFollowStatus.ToString(), domain.DeactiveFollowStatus.ToString()},

		ID:    offset,
		Limit: limit,
	}

	if limit <= 0 || limit > MAX_DEFAULT_LIMIT_REPO {
		queryReq.Limit = MAX_DEFAULT_LIMIT_REPO
	}

	rows, err := r.querier.AllByFollowedUserID(ctx, queryReq)
	if err != nil {
		return domain.Follows{}, errors.New(fmt.Sprintf("querier.AllByFollowedUserID %s", err.Error()))
	}

	follows := domain.NewEmptyFollows(len(rows))
	for i := 0; i < len(rows); i++ {
		follows = append(follows, domain.Follow{
			ID:              rows[i].ID,
			FollowedUserID:  rows[i].FollowedUserID,
			FollowingUserID: rows[i].FollowingUserID,
			Status:          domain.NewFollowStatusFromString(rows[i].Status),

			CreatedDate: &rows[i].CreatedDate,
			UpdatedDate: &rows[i].UpdatedDate.Time,
		})
	}

	return follows, nil
}

func (r *repoManager) Update(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID, status domain.FollowStatus) (domain.Follow, error) {

	row, err := r.querier.UpdateFollow(ctx, followPgQeries.UpdateFollowParams{
		Status:          status.ToString(),
		FollowedUserID:  followedUserID,
		FollowingUserID: followingUserID,
	})
	if err != nil {
		return domain.Follow{}, errors.New(fmt.Sprintf("querier.AllByFollowedUserID %s", err.Error()))
	}

	return domain.Follow{
		ID:              row.ID,
		FollowedUserID:  row.FollowedUserID,
		FollowingUserID: row.FollowingUserID,
		Status:          domain.NewFollowStatusFromString(row.Status),
		CreatedDate:     &row.CreatedDate,
		UpdatedDate:     &row.UpdatedDate.Time,
	}, nil
}

func (r *repoManager) CountByFollowedUserID(ctx context.Context, followedUserID uuid.UUID) (int64, error) {

	count, err := r.querier.CountByFollowedUserID(ctx, followedUserID)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("querier.CountByFollowedUserID %s", err.Error()))
	}

	return count, nil
}

func (r *repoManager) GetOne(ctx context.Context, followedUserID uuid.UUID, followingUserID uuid.UUID) (domain.Follow, error) {

	row, err := r.querier.GetOne(ctx, followPgQeries.GetOneParams{
		FollowedUserID:  followedUserID,
		FollowingUserID: followingUserID,
	})
	if err != nil {
		return domain.Follow{}, errors.New(fmt.Sprintf("querier.CountByFollowedUserID %s", err.Error()))
	}

	return domain.Follow{
		ID:              row.ID,
		FollowedUserID:  row.FollowedUserID,
		FollowingUserID: row.FollowingUserID,
		Status:          domain.NewFollowStatusFromString(row.Status),
		CreatedDate:     &row.CreatedDate,
		UpdatedDate:     &row.UpdatedDate.Time,
	}, nil
}
