package job

import (
	"github.com/Suhaan-Bhandary/go-api-template/internal/job/cronjob"
	"github.com/go-co-op/gocron/v2"
)

func InitializeJobs(scheduler gocron.Scheduler) {
	sayHelloJob := cronjob.NewSayHelloJob("Hi, running Cron Job", scheduler)
	sayHelloJob.Schedule()

	sayHelloDailyJob := cronjob.NewSayHelloDailyJob(scheduler)
	sayHelloDailyJob.Schedule()
}
