package cron

import "github.com/robfig/cron/v3"

// 定时任务的结构体
type Cron struct {
	Func  func() // 定时任务的函数
	Timer string // 定时任务的时间
}

func Start(crons ...Cron) {
	cron2 := cron.New(cron.WithSeconds()) //创建一个cron实例

	for _, v := range crons {
		cron2.AddFunc(v.Timer, v.Func)
	}

	cron2.Start()
}
