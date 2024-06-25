package ctxlogger

import (
	"context"
	"os"

	customcontext "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/environment"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)

	if environment.ENVIRONMENT == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
}

func Info(ctx context.Context, message string, args ...interface{}) {
	rid := customcontext.GetRequestID(ctx)
	log.WithField("rid", rid).Infof(message, args...)
}

func Warn(ctx context.Context, message string, args ...interface{}) {
	rid := customcontext.GetRequestID(ctx)
	log.WithField("rid", rid).Warnf(message, args...)
}

func Error(ctx context.Context, message string, args ...interface{}) {
	rid := customcontext.GetRequestID(ctx)
	log.WithField("rid", rid).Errorf(message, args...)
}

func Debug(ctx context.Context, message string, args ...interface{}) {
	rid := customcontext.GetRequestID(ctx)
	log.WithField("rid", rid).Debugf(message, args...)
}
