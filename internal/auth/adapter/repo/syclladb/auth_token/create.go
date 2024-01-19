package auth_token

import (
	"context"
	"time"

	"github.com/duongnln96/blog-realworld/internal/auth/core/domain"
)

func (r *repoManager) Create(ctx context.Context, model domain.AuthToken) (domain.AuthToken, error) {

	now := time.Now()
	model.CreatedDate = &now
	model.UpdatedDate = &now

	insertModel := AuthToken{
		ID:          GoogleUUIDMarshaler{model.ID},
		UserID:      GoogleUUIDMarshaler{model.UserID},
		DeviceID:    GoogleUUIDMarshaler{model.DeviceID},
		UserAgent:   model.UserAgent,
		JwtToken:    model.JwtToken.ToString(),
		RemoteIP:    model.RemoteIP,
		ExpiredDate: model.ExpiredDate,
		Status:      model.Status.ToString(),

		CreatedDate: model.CreatedDate,
		UpdatedDate: model.UpdatedDate,
	}

	q := r.db.GetSession().Query(authenTokenTable.Insert()).WithContext(ctx).BindStruct(insertModel)
	if err := q.ExecRelease(); err != nil {
		return domain.AuthToken{}, err
	}

	return model, nil

}
