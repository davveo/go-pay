package initialize

import (
	"github.com/davveo/go-pay/initialize/conf"
	"github.com/davveo/go-pay/initialize/cron"
	"github.com/davveo/go-pay/initialize/db"
	"github.com/davveo/go-pay/initialize/log"
	"github.com/davveo/go-pay/initialize/migrate"
	"github.com/davveo/go-pay/initialize/mq"
	"github.com/davveo/go-pay/initialize/redis"
	"github.com/davveo/go-pay/utils/public"
)

func Init() {
	var err error
	// 配置初始化
	err = conf.Init()
	public.CheckError(err)

	// 日志初始化
	err = log.Init()
	public.CheckError(err)

	// Redis初始化
	err = redis.Init()
	public.CheckError(err)

	// 数据库初始化
	err = db.Init()
	public.CheckError(err)

	// mq初始化
	err = mq.Init()
	public.CheckError(err)

	// 数据库健康检测
	err = cron.Init()
	public.CheckError(err)

	// 数据库迁移
	err = migrate.Init()
	public.CheckError(err)

}
