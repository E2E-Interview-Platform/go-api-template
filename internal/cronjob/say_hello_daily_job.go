package cronjob

import (
	"context"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/go-co-op/gocron/v2"
)

const SAY_HELLO_DAILY_CRON_JOB = "SAY_HELLO_DAILY_CRON_JOB"
const SAY_HELLO_DAILY_CRON_JOB_INTERVAL_DAYS = 1

var SayHelloDailyJobTiming = JobTime{
	hours:   0,
	minutes: 0,
	seconds: 0,
}

type SayHelloDailyJob struct {
	CronJob
}

func NewSayHelloDailyJob(scheduler gocron.Scheduler) Job {
	return &SayHelloDailyJob{
		CronJob: CronJob{
			name:      SAY_HELLO_DAILY_CRON_JOB,
			scheduler: scheduler,
		},
	}
}

func (cron *SayHelloDailyJob) Schedule() {
	var err error
	cron.job, err = cron.scheduler.NewJob(
		gocron.DailyJob(
			SAY_HELLO_DAILY_CRON_JOB_INTERVAL_DAYS,
			gocron.NewAtTimes(
				gocron.NewAtTime(
					SayHelloDailyJobTiming.hours,
					SayHelloDailyJobTiming.minutes,
					SayHelloDailyJobTiming.seconds,
				),
			),
		),
		gocron.NewTask(cron.Execute, cron.Task),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	cron.scheduler.Start()

	if err != nil {
		ctxlogger.Warn(context.TODO(), "error occurred while scheduling %s, message %+v", cron.name, err.Error())
	}
}

func (cron *SayHelloDailyJob) Task(ctx context.Context) {
	ctxlogger.Info(ctx, "Hello!!")
}
