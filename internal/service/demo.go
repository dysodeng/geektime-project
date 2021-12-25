package service

import (
	"github.com/dysodeng/project/internal/entity"
	"github.com/dysodeng/project/internal/service/contracts"
	"github.com/dysodeng/project/pkg/db"
	"github.com/pkg/errors"
)

type DemoService struct{}

var _ contracts.DemoServiceInterface = new(DemoService)

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (demoService DemoService) UserInfo(userId uint64) (entity.User, error) {
	if userId == 0 {
		return entity.User{}, errors.New("缺少用户ID")
	}

	var user entity.User
	db.DB().Where("id", userId).First(&user)

	if user.ID == 0 {
		return entity.User{}, errors.New("用户不存在")
	}

	return user, nil
}
