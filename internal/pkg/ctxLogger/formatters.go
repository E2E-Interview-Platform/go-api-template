package ctxlogger

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

// Customer formatter for logrus
type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b strings.Builder

	level := entry.Level
	coloredMessage := helpers.ColorizeForLevel(strings.ToUpper(level.String()), level)
	timestamp := entry.Time.Format(time.RFC3339)

	b.WriteString(fmt.Sprintf("\n[%s:%s]\n%s\n", coloredMessage, timestamp, strings.TrimSpace(entry.Message)))

	// sorted slice of keys for consistent ordering
	keys := make([]string, 0, len(entry.Data))
	for key := range entry.Data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		b.WriteString(fmt.Sprintf("[%s=%v] ", helpers.ColorizeForLevel(key, level), entry.Data[key]))
	}

	b.WriteString("\n")

	return []byte(b.String()), nil
}
