package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var settings C
var v *viper.Viper

func Init() error {
	v = viper.New()
	v.SetConfigName("pay")
	v.SetConfigType("yaml")
	v.AddConfigPath("./conf")

	if err := v.ReadInConfig(); err != nil {
		logrus.Fatal(err)
		return err
	}
	if err := v.Unmarshal(&settings); err != nil {
		logrus.Fatal(err)
		return err
	} else {
		v.WatchConfig()
		logrus.Info("配置初始化成功")
	}
	return nil
}

func GetSettings() C {
	return settings
}
