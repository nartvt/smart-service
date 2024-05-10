package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/nartvt/go-core/uerror"
	"github.com/nartvt/go-core/utils"
	v1 "github.com/nartvt/smart-service/api/smart/v1"
	"github.com/nartvt/smart-service/internal/module"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user module.User) error
	GetUserByUserName(ctx context.Context, userName string) (module.User, error)
	GetUserByUserNameOrEmail(ctx context.Context, userName, email string) (module.User, error)
}

type UserBiz struct {
	repo IUserRepo
	log  *log.Helper
}

func NewUserBiz(repo IUserRepo) *UserBiz {
	return &UserBiz{
		repo: repo,
		log:  log.NewHelper(log.With(log.DefaultLogger, "module", "biz/user")),
	}
}

func (u *UserBiz) CreateUser(ctx context.Context, user *v1.UserRegisterRequest) error {
	u.log.WithContext(ctx).Infof("create user: %v", user)
	return u.repo.CreateUser(ctx, mapUserRegisterToModule(user))
}

func mapUserRegisterToModule(user *v1.UserRegisterRequest) module.User {
	hashPassword, _ := utils.HashPassword(user.Password)
	return module.User{
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		UserName:     user.UserName,
		HashPassword: hashPassword,
	}
}

func (u *UserBiz) Login(ctx context.Context, userRq *v1.UserLoginRequest) (*v1.UserLoginResponse, error) {
	if userRq == nil {
		return nil, uerror.UnAuthorizeError("invalid user")
	}

	user, err := u.repo.GetUserByUserNameOrEmail(ctx, userRq.UserName, userRq.Email)
	if err != nil {
		return nil, uerror.InteralServerError(err.Error())
	}
	hashPassword, _ := utils.HashPassword(userRq.Password)
	if hashPassword != user.HashPassword {
		return nil, uerror.UnAuthorizeError("invalid user")
	}
	return &v1.UserLoginResponse{}, nil
}
