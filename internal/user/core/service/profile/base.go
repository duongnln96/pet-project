package profile

import (
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/google/wire"
)

type service struct {
	config *config.Configs

	followRepo port.FollowRepoI
	userRepo   port.UserRepoI
}

var _ port.FollowServiceI = (*service)(nil)

var ServiceSet = wire.NewSet(NewService)

func NewService(
	config *config.Configs,

	followRepoInstance port.FollowRepoI,
	userRepoInstance port.UserRepoI,
) port.FollowServiceI {

	return &service{
		config:     config,
		followRepo: followRepoInstance,
		userRepo:   userRepoInstance,
	}
}
