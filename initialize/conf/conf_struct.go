package conf

import "time"

type C struct {
	Mysql  Mysql
	Redis  Redis
	Server Server
	Mq     Mq
}

type Mysql struct {
	DbName      string
	Host        string
	Port        string
	User        string
	Password    string
	MaxIdleConn int
	MaxOpenConn int
	ShowSQL     bool
}

type Redis struct {
	Host        string
	Port        int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Server struct {
	Port int
}

type Mq struct {
	Host     string
	Port     int
	User     string
	Password string
}
