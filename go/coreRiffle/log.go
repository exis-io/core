package coreRiffle

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

const (
	levelWarn  int = 1
	levelInfo  int = 2
	levelDebug int = 3
)

var logLevel int = 0

func Debug(format string, a ...interface{}) {
	if logLevel >= levelDebug {
		out(fmt.Sprintf(format, a...), color.Blue)
	}
}

func Info(format string, a ...interface{}) {
	if logLevel >= levelInfo {
		out(fmt.Sprintf(format, a...), color.Green)
	}
}

func Warn(format string, a ...interface{}) {
	if logLevel >= levelWarn {
		out(fmt.Sprintf(format, a...), color.Yellow)
	}
}

func SetLogging(level int) {
	logLevel = level
}

func out(mess string, printer func(string, ...interface{})) {
	printer("[%s] %s", trace(), mess)
}

func trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(4, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	parts := strings.Split(file, "/")

	if len(parts) > 3 {
		last := parts[len(parts)-2:]
		file = strings.Join(last, ".")
	}

	return fmt.Sprintf("%s:%d", strings.TrimSuffix(file, ".go"), line)
}

// var format = logging.MustStringFormatter(
// 	// "%{color}[%{time:2006-01-02 15:04:05.000} %{longfunc}] %{message}%{color:reset}",
// 	"[%{color}%{longfunc}]  %{message}",
// )

// func InitLogger() {
// 	// For demo purposes, create two backend for os.Stderr.
// 	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
// 	formatter := logging.NewBackendFormatter(backend1, format)
// 	backend1Leveled := logging.AddModuleLevel(backend1)

// 	if os.Getenv("DEBUG") != "" {
// 		backend1Leveled.SetLevel(logging.DEBUG, "")
// 	} else {
// 		backend1Leveled.SetLevel(logging.CRITICAL, "")
// 	}

// 	logging.SetBackend(backend1Leveled, formatter)

// 	// out.Debug("debug")
// 	// out.Info("info")
// 	// out.Notice("notice")
// 	// out.Warning("warning")
// 	// out.Error("err")
// 	// out.Critical("crit")
// 	Log.Debug("Logger initialized")

// 	trace()
// }

// func logErr(err error) error {
// 	if err == nil {
// 		return nil
// 	}

// 	return err
// }

// var Log = logging.MustGetLogger("example")
