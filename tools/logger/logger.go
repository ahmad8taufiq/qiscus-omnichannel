package logger

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()
type ColorFormatter struct{}

func InitLogger() {
	Logger.SetFormatter(&ColorFormatter{}) // your custom formatter
	Logger.SetLevel(logrus.DebugLevel)
}

func (f *ColorFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer

	// Format timestamp
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	levelColor, levelLabel := getColorAndLabel(entry.Level)

	// Main log line
	b.WriteString(fmt.Sprintf("%s %s ▶ %s\n",
		levelColor(fmt.Sprintf("[%s]", levelLabel)),
		color.New(color.FgWhite).Sprint(timestamp),
		entry.Message,
	))

	// Multiline key:value fields
	for key, val := range entry.Data {
		valStr := formatValue(val)
		b.WriteString(fmt.Sprintf("    %s: %s\n",
			color.New(color.FgCyan).Sprint(key),
			valStr,
		))
	}

	return b.Bytes(), nil
}

func getColorAndLabel(level logrus.Level) (func(a ...interface{}) string, string) {
	switch level {
	case logrus.DebugLevel:
		return color.New(color.FgHiBlue).SprintFunc(), "DEBUG"
	case logrus.InfoLevel:
		return color.New(color.FgGreen).SprintFunc(), "INFO"
	case logrus.WarnLevel:
		return color.New(color.FgYellow).SprintFunc(), "WARN"
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return color.New(color.FgRed).SprintFunc(), "ERROR"
	default:
		return color.New(color.FgWhite).SprintFunc(), "LOG"
	}
}

func formatValue(v interface{}) string {
	switch val := v.(type) {
	case string:
		if json.Valid([]byte(val)) {
			var pretty bytes.Buffer
			if err := json.Indent(&pretty, []byte(val), "", "    "); err == nil {
				return pretty.String()
			}
		}
		return val
	default:
		jsonVal, _ := json.MarshalIndent(val, "", "    ")
		return string(jsonVal)
	}
}
