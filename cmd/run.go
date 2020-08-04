package cmd

import (
	"strconv"
	"time"

	"github.com/davveo/go-pay/initialize"
	C "github.com/davveo/go-pay/initialize/conf"
	"github.com/davveo/go-pay/router"
	"github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
)

func init() {
	// 初始化操作
	initialize.Init()
}

func Runserver() error {
	// 启动服务
	logrus.Println("Starting goPay server on port ",
		C.GetSettings().Server.Port)
	graceful.Run(":"+strconv.Itoa(C.GetSettings().Server.Port),
		5*time.Second, router.SetupRouter())
	return nil
}
