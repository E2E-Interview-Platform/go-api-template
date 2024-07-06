package cronjob

import (
	"context"
	"time"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/go-co-op/gocron/v2"
)

const SAY_HELLO_CRON_JOB = "SAY_HELLO_CRON_JOB"
const SAY_HELLO_CRON_JOB_INTERVAL_MINUTES = 1

type SayHelloJob struct {
	CronJob
	message string
}

func NewSayHelloJob(
	message string,
	scheduler gocron.Scheduler,
) Job {
	return &SayHelloJob{
		message: message,
		CronJob: CronJob{
			name:      SAY_HELLO_CRON_JOB,
			scheduler: scheduler,
		},
	}
}

func (cron *SayHelloJob) Schedule() {
	var err error
	cron.job, err = cron.scheduler.NewJob(
		gocron.DurationJob(SAY_HELLO_CRON_JOB_INTERVAL_MINUTES*time.Minute),
		gocron.NewTask(cron.Execute, cron.Task),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	cron.scheduler.Start()

	if err != nil {
		ctxlogger.Warn(context.TODO(), "error occurred while scheduling %s, message %+v", cron.name, err.Error())
	}
}

func (cron *SayHelloJob) Task(ctx context.Context) {
	ctxlogger.Info(ctx, cron.message)
}
