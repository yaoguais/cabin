package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *Logger

const (
	Ldebug = iota
	Linfo
	Lwarn
	Lerror
	Lpanic
	Lfatal
)

func init() {
	logger = New()
}

type Logger struct {
	config zap.Config
	logger *zap.Logger
}

func New() *Logger {
	config := zap.NewProductionConfig()
	logger, err := config.Build()
	if err != nil {
		panic("new logger failed")
	}

	return &Logger{
		config: config,
		logger: logger,
	}
}
func SetOutputLevel(lvl int) { logger.SetOutputLevel(lvl) }
func (l *Logger) SetOutputLevel(lvl int) {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.Level(int8(lvl - 1)))
	l.config.Level = level
	log, err := l.config.Build()
	if err != nil {
		panic("set log output level failed")
	}
	l.logger = log
}

func SetOutputPath(path string) { logger.SetOutputPath(path) }
func (l *Logger) SetOutputPath(path string) {
	l.config.OutputPaths = []string{path}
	l.config.ErrorOutputPaths = []string{path}
	log, err := l.config.Build()
	if err != nil {
		panic("set log output level failed")
	}
	l.logger = log
}

// -----------------------------------------

func Printf(format string, v ...interface{}) { logger.Printf(format, v...) }
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, v...))
}

func Print(v ...interface{}) { logger.Print(v...) }
func (l *Logger) Print(v ...interface{}) {
	l.logger.Info(fmt.Sprint(v...))
}

func Println(v ...interface{}) { logger.Println(v...) }
func (l *Logger) Println(v ...interface{}) {
	l.logger.Info(fmt.Sprintln(v...))
}

// -----------------------------------------
func Debugf(format string, v ...interface{}) { logger.Debugf(format, v...) }
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) { logger.Debug(v...) }
func (l *Logger) Debug(v ...interface{}) {
	l.logger.Debug(fmt.Sprint(v...))
}

// -----------------------------------------

func Infof(format string, v ...interface{}) { logger.Infof(format, v...) }
func (l *Logger) Infof(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) { logger.Info(v...) }
func (l *Logger) Info(v ...interface{}) {
	l.logger.Info(fmt.Sprint(v...))
}

// -----------------------------------------
func Warnf(format string, v ...interface{}) { logger.Warnf(format, v...) }
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) { logger.Warn(v...) }
func (l *Logger) Warn(v ...interface{}) {
	l.logger.Warn(fmt.Sprint(v...))
}

// -----------------------------------------
func Errorf(format string, v ...interface{}) { logger.Errorf(format, v...) }
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) { logger.Error(v...) }
func (l *Logger) Error(v ...interface{}) {
	l.logger.Error(fmt.Sprint(v...))
}

// -----------------------------------------
func Fatal(v ...interface{}) { logger.Fatal(v...) }
func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) { logger.Fatalf(format, v...) }
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(format, v...))
}

func Fatalln(v ...interface{}) { logger.Fatalln(v...) }
func (l *Logger) Fatalln(v ...interface{}) {
	l.logger.Fatal(fmt.Sprintln(v...))
}

// -----------------------------------------
func Panic(v ...interface{}) { logger.Panic(v...) }
func (l *Logger) Panic(v ...interface{}) {
	l.logger.Panic(fmt.Sprint(v...))
}

func Panicf(format string, v ...interface{}) { logger.Panicf(format, v...) }
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.logger.Panic(fmt.Sprintf(format, v...))
}

func Panicln(v ...interface{}) { logger.Panicln(v...) }
func (l *Logger) Panicln(v ...interface{}) {
	l.logger.Panic(fmt.Sprintln(v...))
}
