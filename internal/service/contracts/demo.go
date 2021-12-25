package contracts

import "github.com/dysodeng/project/internal/entity"

type DemoServiceInterface interface {
	UserInfo(userId uint64) (entity.User, error)
}
