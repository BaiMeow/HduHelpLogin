package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

const (
	greenGround   = "\033[97;42m"
	whiteGround   = "\033[90;47m"
	yellowGround  = "\033[90;43m"
	redGround     = "\033[97;41m"
	blueGround    = "\033[97;44m"
	magentaGround = "\033[97;45m"
	cyanGround    = "\033[97;46m"
	resetGround   = "\033[0m"
)

type Formatter struct{}

var defaultFormatter = logrus.TextFormatter{ForceColors: true}

func (f Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	Type, ok := entry.Data["type"].(string)
	if !ok {
		return defaultFormatter.Format(entry)
	}
	levelColor := levelColor(entry.Level)
	levelText := strings.ToUpper(entry.Level.String())
	TypeText := strings.ToUpper(Type)
	b := new(bytes.Buffer)
	switch Type {
	case "gorm":
		fmt.Fprintf(b, "\x1b[%dm%-11s\x1b[0m%s", levelColor, TypeText+"-"+levelText, entry.Time.Format("2006/01/02 - 15:04:05"))
		fmt.Fprintf(b, "|%srows:%s %s| %s |%s |%s\n%-11s%s",
			blueGround, safeGetValue[string](entry, "rows"), resetGround,
			safeGetValue[string](entry, "duration"),
			safeGetValue[string](entry, "fileWithLineNum"),
			FormatTraceIdIfExist(entry),
			"",
			safeGetValue[string](entry, "sql"))
	case "gin":
		timeStamp := safeGetValue[string](entry, "timeStamp")
		fmt.Fprintf(b, "\x1b[%dm%-11s\x1b[0m%s", levelColor, TypeText+"-"+levelText, timeStamp)
		method := safeGetValue[string](entry, "method")
		statusCode := safeGetValue[int](entry, "statusCode")
		statusColor := StatusCodeColor(statusCode)
		methodColor := MethodColor(method)
		resetColor := resetGround
		Latency := safeGetValue[time.Duration](entry, "Latency")
		if Latency > time.Minute {
			Latency = Latency.Truncate(time.Second)
		}
		fmt.Fprintf(b, "|%s %3d %s| %13v | %15s |%s %-7s %s %#v |%s",
			statusColor, statusCode, resetColor,
			Latency,
			safeGetValue[string](entry, "clientIP"),
			methodColor, method, resetColor,
			safeGetValue[string](entry, "path"),
			FormatTraceIdIfExist(entry),
		)
		if entry.Message != "" {
			fmt.Fprintf(b, "\n%-11s%s", "", entry.Message)
		}
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func StatusCodeColor(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return greenGround
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return whiteGround
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellowGround
	default:
		return redGround
	}
}

func MethodColor(method string) string {
	switch method {
	case http.MethodGet:
		return blueGround
	case http.MethodPost:
		return cyanGround
	case http.MethodPut:
		return yellowGround
	case http.MethodDelete:
		return redGround
	case http.MethodPatch:
		return greenGround
	case http.MethodHead:
		return magentaGround
	case http.MethodOptions:
		return whiteGround
	default:
		return resetGround
	}
}

func levelColor(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return gray
	case logrus.WarnLevel:
		return yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return red
	case logrus.InfoLevel:
		return blue
	default:
		return blue
	}
}

func safeGetValue[T any](entry *logrus.Entry, key string) T {
	V, _ := entry.Data[key].(T)
	return V
}

func FormatTraceIdIfExist(entry *logrus.Entry) string {
	uu, ok := entry.Data["traceId"]
	if !ok {
		return ""
	}
	return fmt.Sprint(uu)
}
