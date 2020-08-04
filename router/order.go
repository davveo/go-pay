package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrderRouter(router *gin.RouterGroup) {
	// 前端路由
	orderRouter := router.Group("order")
	{
		orderRouter.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"success": "ok",
			})
		})
	}

}
