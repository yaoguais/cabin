package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var (
	uploader *Uploader
	slackbot *SlackBot
)

func main() {
	configFile := flag.String("config", "config.toml", "config file")
	pprof := flag.String("pprof", "", "pprof address e,g. 127.0.0.1:6060")
	flag.Parse()

	if *pprof != "" {
		go http.ListenAndServe(*pprof, nil)
	}

	c, err := LoadConfig(*configFile)
	if err != nil {
		logrus.WithError(err).Fatal("load config")
	}

	setupLogger(c)
	start(c)
}

func setupLogger(c *Config) {
	l, err := logrus.ParseLevel(c.Logger.Level)
	if err != nil {
		logrus.WithError(err).Fatal("illegal logger level")
	}
	logrus.SetLevel(l)

	format := "2006-01-02T15:04:05.000-07:00"
	fieldMap := logrus.FieldMap{
		logrus.FieldKeyTime:  "@time",
		logrus.FieldKeyLevel: "@level",
		logrus.FieldKeyMsg:   "@message",
	}

	switch c.Logger.Formatter {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: format,
		})
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: format,
			FieldMap:        fieldMap,
		})
	default:
		logrus.WithField("formatter", c.Logger.Formatter).Fatal("formatter should be text or json")
	}

	if c.Logger.File != "" {
		logrus.SetOutput(&lumberjack.Logger{
			Filename:   c.Logger.File,
			MaxSize:    c.Logger.MaxSize,
			MaxBackups: c.Logger.MaxBackups,
			MaxAge:     c.Logger.MaxAge,
			Compress:   c.Logger.Compress,
		})
	}
}

func start(c *Config) {
	serveErr := make(chan error, 10)
	uploader, err := NewUploader(c.Uploader)
	if err != nil {
		logrus.WithError(err).Fatal("create uploader")
	}
	uploader.Serve(serveErr)
	logrus.WithField("address", fmt.Sprintf("%s:%d", c.Uploader.Host, c.Uploader.Port)).Info("uploader listen on")

	slackbot, err = NewSlackBot(c.SlackBot)
	if err != nil {
		logrus.WithError(err).Fatal("create slackbot")
	}
	slackbot.Serve(serveErr)
	logrus.WithField("address", fmt.Sprintf("%s:%d", c.SlackBot.Host, c.SlackBot.Port)).Info("slackbot listen on")

	sg := make(chan os.Signal, 1)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	select {
	case err = <-serveErr:
		logrus.WithError(err).Fatal("serve error")
	case s := <-sg:
		logrus.WithField("signal", s.String()).Info("receive signal")
		if err = uploader.Shutdown(); err != nil {
			logrus.WithError(err).Error("shutdown uploader server")
		}
		if err = slackbot.Shutdown(); err != nil {
			logrus.WithError(err).Error("shutdown slackbot server")
		}
	}
}
