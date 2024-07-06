package cronjob

import (
	"context"
	"fmt"
	"time"

	customcontext "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/go-co-op/gocron/v2"
)

type Job interface {
	// Schedules the cron job
	Schedule()
}

type CronJob struct {
	name      string
	scheduler gocron.Scheduler
	job       gocron.Job
}

type JobTime struct {
	hours   uint
	minutes uint
	seconds uint
}

// Execute function setups the basic config and runs the cron task
func (cron *CronJob) Execute(task func(context.Context)) {
	rid := fmt.Sprintf("[CRON]-%s", cron.name)

	ctx := context.Background()
	ctx = customcontext.SetRequestID(ctx, rid)

	startTime := time.Now()
	ctxlogger.Info(ctx, "cron job Started at %s", startTime.Format("2006-01-02 15:04:05"))
	defer func() {
		ctxlogger.Info(ctx, "cron job done %s, took: %v", cron.name, time.Since(startTime))
	}()

	// Channel to check if signal task is completed
	taskCompletedSignalChan := make(chan struct{})

	// Executing cron job in separate go routine
	go func() {
		defer func() {
			taskCompletedSignalChan <- struct{}{}
		}()

		task(ctx)
	}()

	// Blocking till task completes
	<-taskCompletedSignalChan
}
