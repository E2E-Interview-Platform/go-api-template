package helpers

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	AnsiReset     = 0
	AnsiRed       = 31
	AnsiHiRed     = 91
	AnsiGreen     = 32
	AnsiHiGreen   = 92
	AnsiYellow    = 33
	AnsiHiYellow  = 93
	AnsiBlue      = 34
	AnsiHiBlue    = 94
	AnsiMagenta   = 35
	AnsiHiMagenta = 95
	AnsiCyan      = 36
	AnsiHiCyan    = 96
	AnsiWhite     = 37
	AnsiHiWhite   = 97
)

func ColorizeForLevel(message string, level logrus.Level) string {
	var levelColor int

	switch level {
	case logrus.DebugLevel:
		levelColor = AnsiMagenta
	case logrus.WarnLevel:
		levelColor = AnsiYellow
	case logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel:
		levelColor = AnsiRed
	case logrus.InfoLevel:
		levelColor = AnsiCyan
	default:
		levelColor = 0
	}

	return Colorize(message, levelColor)
}

func Colorize(message string, levelColor int) string {
	if levelColor == 0 {
		return message
	}

	return "\033[" + strconv.Itoa(levelColor) + "m" + message + "\033[0m"
}
