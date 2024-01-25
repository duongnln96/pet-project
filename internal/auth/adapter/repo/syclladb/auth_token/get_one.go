package auth_token

import (
	"context"
	"errors"

	"github.com/duongnln96/blog-realworld/internal/auth/core/domain"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func (r *repoManager) GetOneByPrimary(ctx context.Context, tokenID uuid.UUID) (domain.AuthToken, error) {

	row := AuthTokenRow{}

	if err := r.db.GetSession().Query(authenTokenTable.Select()).WithContext(ctx).Bind(GoogleUUIDMarshaler{tokenID}).Get(&row); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return domain.AuthToken{}, nil
		}
		return domain.AuthToken{}, err
	}

	return domain.AuthToken{
		ID:          row.ID.UUID,
		UserID:      row.UserID.UUID,
		DeviceID:    row.DeviceID.UUID,
		UserAgent:   row.UserAgent,
		JwtToken:    domain.NewJwtAuthTokenFromStr(row.JwtToken),
		RemoteIP:    row.RemoteIP,
		ExpiredDate: row.ExpiredDate,
		Status:      domain.NewAuthTokenStatusFromStr(row.Status),

		CreatedDate: row.CreatedDate,
		UpdatedDate: row.UpdatedDate,
	}, nil
}
