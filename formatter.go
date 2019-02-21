package cronenberg

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Formatter struct{}

func (formatter *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	entry.Message = strings.TrimSuffix(entry.Message, "\n")

	fmt.Fprintf(
		b,
		entryFormat,
		levelText,
		timestamp(entry.Time.Sub(baseTimestamp)),
		entry.Data["context"],
		entry.Message,
	)

	return b.Bytes(), nil
}

func timestamp(t time.Duration) string {
	hours := t / time.Hour
	t = t - (hours * time.Hour)

	minutes := t / time.Minute
	t = t - (minutes * time.Minute)

	seconds := t / time.Second

	return fmt.Sprintf("%02dh%02dm%02ds", int(hours), int(minutes), int(seconds))
}

var baseTimestamp time.Time

var entryFormat = "%s[%s](%s) %s\n"

func init() {
	baseTimestamp = time.Now()
}
