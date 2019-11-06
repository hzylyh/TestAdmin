package conf

import "github.com/robfig/cron"

var (
	JobList = make(map[string]*cron.Cron)
)
