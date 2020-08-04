package cron

import (
	"github.com/robfig/cron/v3"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func CheckMysqlHealth() {

}

func Init() error {
	c := newWithSeconds()
	_, err := c.AddFunc("*/5 * * * * ?", CheckMysqlHealth)
	if err != nil {
		return err
	} else {
		c.Start()
	}

	return nil
}
