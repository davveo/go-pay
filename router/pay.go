package router

import (
	"github.com/gin-gonic/gin"
)

func OrderRouter(router *gin.RouterGroup) {
	// 前端路由
	orderRouter := router.Group("")
	{
		// 支付
		orderRouter.POST("gopay", func(context *gin.Context) {

		})
		// 退款
		orderRouter.POST("refund", func(context *gin.Context) {

		})
		// 异步通知
		orderRouter.POST("notify/paytype/payno", func(context *gin.Context) {

		})
		// 同步通知
		orderRouter.POST("return/paytype/payno", func(context *gin.Context) {

		})
		// 交易查询
		orderRouter.GET("query/trade", func(context *gin.Context) {

		})
		// 退款查询
		orderRouter.GET("query/refund", func(context *gin.Context) {

		})

		// 账单获取
		orderRouter.GET("query/bill", func(context *gin.Context) {

		})

		// 结算明细
		orderRouter.GET("query/settle", func(context *gin.Context) {

		})

	}

}
