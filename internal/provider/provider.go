package provider

import (
	"fmt"
	"github.com/dysodeng/project/internal/service"
	"github.com/dysodeng/project/internal/service/contracts"
	"github.com/goava/di"
	"github.com/pkg/errors"
)

// ServiceProvider 服务容器初始化
func ServiceProvider() {
	var err error
	service.Container, err = di.New(
		// 服务注入服务容器
		di.Provide(service.NewDemoService, di.As(new(contracts.DemoServiceInterface))),
	)
	if err != nil {
		err = errors.Wrap(err, "service provider error.")
		fmt.Printf("%+v\n", err)
		panic(err)
	}
}
