package http

import (
	"github.com/dysodeng/project/internal/service"
	"github.com/dysodeng/project/internal/service/contracts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.GET("user", func(ctx *gin.Context) {
		var body Body
		_ = ctx.ShouldBindJSON(&body)

		var demoService contracts.DemoServiceInterface
		_ = service.Resolve(&demoService)

		user, err := demoService.UserInfo(body.UserId)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user": map[string]interface{}{
				"id":       user.ID,
				"username": user.Username,
				"nickname": user.Nickname,
			},
		})
	})

	return router
}

type Body struct {
	UserId uint64 `form:"user_id" json:"user_id"`
}
