package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/nlopes/slack"
)

// SlackBotConfig config for uploader
type SlackBotConfig struct {
	Host   string `json:"host" toml:"host" yaml:"host"`
	Port   int    `json:"port" toml:"port" yaml:"port"`
	Path   string `json:"path" toml:"path" yaml:"path"`
	Token  string `json:"token" toml:"token" yaml:"token"`
	Secret string `json:"secret" toml:"secret" yaml:"secret"`
	RTM    bool   `json:"rtm" toml:"rtm" yaml:"rtm"`
}

// SlackBot slackbot
type SlackBot struct {
	c      *SlackBotConfig
	s      *http.Server
	client *slack.Client
	rtm    *slack.RTM
}

// NewUploader create new instance for SlackBot
func NewSlackBot(c SlackBotConfig) (*SlackBot, error) {
	client := slack.New(c.Token)
	_, err := client.AuthTest()
	if err != nil {
		return nil, err
	}
	s := &SlackBot{
		c:      &c,
		client: client,
	}

	if c.RTM {
		if err = s.SetupRTM(); err != nil {
			return nil, err
		}
	}

	return s, nil
}

// Serve set up server for upload http request
func (s *SlackBot) Serve(c chan error) {
	addr := fmt.Sprintf("%s:%d", s.c.Host, s.c.Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		c <- err
		return
	}

	s.s = &http.Server{}
	mux := http.NewServeMux()
	mux.HandleFunc(s.c.Path, s.send)
	s.s.Handler = loggerMiddleware(mux)

	go func() {
		err = s.s.Serve(ln)
		if err != nil {
			c <- err
		}
	}()
}

// Shutdown gracefully close connections and then shutdown
func (s *SlackBot) Shutdown() error {
	if s.s != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		return s.s.Shutdown(ctx)
	}

	return s.ShutdownRTM()
}

func (s *SlackBot) Send(channel, content string) error {
	_, _, err := s.client.PostMessage(channel, content, slack.PostMessageParameters{})
	return err
}

func (s *SlackBot) send(w http.ResponseWriter, r *http.Request) {
	if s.c.Secret != r.URL.Query().Get("secret") {
		Render(w, r, []byte("permission deny\n"), http.StatusUnauthorized)
		return
	}

	q := r.URL.Query()
	channel := q.Get("channel")
	content := q.Get("content")

	if channel == "" {
		Render(w, r, []byte("channel should't be empty"), http.StatusBadRequest)
		return
	}
	if content == "" {
		Render(w, r, []byte("content should't be empty"), http.StatusBadRequest)
		return
	}

	err := s.Send(channel, content)
	if err != nil {
		Render(w, r, []byte(fmt.Sprintf("send failed, err %v", err)), http.StatusInternalServerError)
		return
	}

	Render(w, r, []byte("send success"), http.StatusOK)
}
