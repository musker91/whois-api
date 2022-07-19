package logger

import (
	"os"
	"path"
	"time"
	"whois-api/configer"
	"whois-api/utils"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	Echo *logrus.Logger
)

func InitialLogger() (err error) {
	Echo = logrus.New()

	// log filename and line number
	//Echo.SetReportCaller(true)

	switch configer.Configer.Serve.LogType {
	case "json":
		Echo.SetFormatter(&logrus.JSONFormatter{})
	default:
		Echo.SetFormatter(&logrus.TextFormatter{})
	}

	if configer.Configer.AppMode == "development" {
		Echo.SetLevel(logrus.DebugLevel)
		Echo.Out = os.Stdout
		return
	}

	switch configer.Configer.Serve.LogLevel {
	case "info":
		Echo.SetLevel(logrus.InfoLevel)
	case "warn":
		Echo.SetLevel(logrus.WarnLevel)
	case "debug":
		Echo.SetLevel(logrus.DebugLevel)
	case "error":
		Echo.SetLevel(logrus.ErrorLevel)
	case "panic":
		Echo.SetLevel(logrus.PanicLevel)
	case "fatal":
		Echo.SetLevel(logrus.FatalLevel)
	default:
		Echo.SetLevel(logrus.DebugLevel)
	}

	switch configer.Configer.Serve.LogOutPath {
	case "file":
		logFileName := path.Join(utils.GetRootPath(), "logs", "whois.log")

		logOut, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
		if err != nil {
			break
		}

		Echo.Out = logOut

		var logSaveDay = time.Duration(configer.Configer.Serve.LogSaveDays)
		var logSplitTime = time.Duration(configer.Configer.Serve.LogSplitTime)

		logWriter, err := rotatelogs.New(
			logFileName+".%Y-%m-%d-%H-%M.log",
			rotatelogs.WithLinkName(logFileName),
			rotatelogs.WithMaxAge(logSaveDay*24*time.Hour),
			rotatelogs.WithRotationTime(logSplitTime*time.Hour),
		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
		Echo.AddHook(lfHook)
	default:
		Echo.Out = os.Stdout
	}
	return
}
