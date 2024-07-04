package ctxlogger

import (
	"context"
	"fmt"
	"os"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constants"
	customcontext "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/environment"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	log.SetOutput(os.Stdout)
	if logDir := os.Getenv(environment.LOG_DIR_KEY); logDir != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/app.log", logDir),
			MaxSize:    constants.MAX_LOG_SIZE_MB,
			MaxBackups: constants.MAX_LOG_BACKUPS,
			Compress:   true,
		})
	}

	if os.Getenv(environment.LOG_FORMAT_KEY) == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetLevel(log.DebugLevel)
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
