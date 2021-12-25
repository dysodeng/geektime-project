package service

import (
	"context"
	"github.com/dysodeng/project/api/proto"
	"github.com/dysodeng/project/internal/service"
	"github.com/dysodeng/project/internal/service/contracts"
	"github.com/pkg/errors"
)

type DemoService struct{}

func (ds *DemoService) UserInfo(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {

	var demoService contracts.DemoServiceInterface
	_ = service.Resolve(&demoService)

	user, err := demoService.UserInfo(req.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "用户查询失败")
	}

	return &proto.UserResponse{
		UserId:   user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
	}, nil
}
