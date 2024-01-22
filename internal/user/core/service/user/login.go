package user

import (
	"context"
	"fmt"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/pkg/utils"
	"github.com/duongnln96/blog-realworld/internal/user/core/domain"
	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/google/uuid"
)

func (s *service) LogIn(ctx context.Context, req *port.LoginUserRequest) (*port.LoginUserResponse, error) {

	err := s.validateEmail(req.Email)
	if err != nil {
		return nil, err
	}

	secretKey, ok := s.config.Other.Get("password_secret_key").(string)
	if !ok {
		return nil, serror.NewSystemSError("cannot get password secret key")
	}
	reqPasswordHashed, err := utils.HashPassword(req.Password, secretKey)
	if err != nil {
		return nil, serror.NewSystemSError(err.Error())
	}

	user, err := s.userRepo.GetOneByEmail(ctx, req.Email)
	if err != nil {
		return nil, serror.NewSystemSError(err.Error())
	}
	if !user.IsExist() {
		return nil, serror.NewSError(domain.NotFoundErrUser, "user not found")
	}
	if user.Password != reqPasswordHashed {
		return nil, serror.NewSError(domain.PasswordInvalidErrUser, "email or password is invalid")
	}

	userAgent, ok := ctx.Value("user-agent").(string)
	if !ok {
		return nil, serror.NewSError(domain.LoginInfoInvalidErrUser, "user-agent invalid parsing")
	}

	deviceIDStr, ok := ctx.Value("device-id").(string)
	if !ok {
		return nil, serror.NewSError(domain.LoginInfoInvalidErrUser, "device-id invalid parsing")
	}
	deviceID, err := uuid.Parse(deviceIDStr)
	if err != nil {
		return nil, serror.NewSError(domain.LoginInfoInvalidErrUser, "device-id invalid parsing")
	}

	remoteIP, ok := ctx.Value("remote-ip").(string)
	if !ok {
		return nil, serror.NewSError(domain.LoginInfoInvalidErrUser, "remote-ip invalid parsing")
	}

	token, err := s.authTokenDomain.GenAuthToken(ctx, &port.GenAuthTokenRequest{
		UserID:    user.ID,
		DeviceID:  deviceID,
		UserAgent: userAgent,
		RemoteIP:  remoteIP,
	})
	if err != nil {
		return nil, serror.NewSystemSError(fmt.Sprintf("authTokenDomain.GenAuthToken %s", err.Error()))
	}

	return &port.LoginUserResponse{
		JwtToken: token.JwtToken,
	}, nil
}
