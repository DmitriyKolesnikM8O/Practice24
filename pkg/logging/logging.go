package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// писать хуки, вне зависимости от уровня: info, warning, debug и так далее
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

// возвращает левелы из хука
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// структура, представляющая собой 1 запись со всеми даннымми
var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithFeild(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

// кастом логгер - записываем все в файлик
// Когда с маленькой буквы - метод вызывается сам всегда, если этот package используется (вызывается до main)
func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// чтобы по умолчанию логи никуда не уходили
	l.SetOutput(io.Discard)

	// записываем логи в all.log и в stdout; logrus.AllLevels - массив всех доступных левелов
	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	// выставляем уровень минимальный, чтобы было видно вообще все
	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
