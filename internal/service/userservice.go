package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/nartvt/smart-service/api/smart/v1"
	"github.com/nartvt/smart-service/internal/biz"
)

type UserServiceService struct {
	log *log.Helper
	pb.UnimplementedUserServiceServer

	userBiz *biz.UserBiz
}

func NewUserServiceService(userBiz *biz.UserBiz) *UserServiceService {
	return &UserServiceService{
		log:     log.NewHelper(log.DefaultLogger),
		userBiz: userBiz,
	}
}

func (s *UserServiceService) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	return &pb.UserRegisterResponse{}, nil
}
func (s *UserServiceService) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return &pb.UserLoginResponse{}, nil
}
