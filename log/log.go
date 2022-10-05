package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func Init() {
	Logger.SetFormatter(Formatter{})

	writer, err := rotatelogs.New("logs/%Y%m%d.log")
	if err != nil {
		Logger.Fatalln(err)
	}
	Logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
		}, &logrus.JSONFormatter{},
	))
}
