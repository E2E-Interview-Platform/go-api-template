package cronjob

import "github.com/go-co-op/gocron/v2"

func InitializeJobs(scheduler gocron.Scheduler) {
	sayHelloJob := NewSayHelloJob("Hi, running Cron Job", scheduler)
	sayHelloJob.Schedule()

	sayHelloDailyJob := NewSayHelloDailyJob(scheduler)
	sayHelloDailyJob.Schedule()
}
