package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

// UploaderConfig config for uploader
type UploaderConfig struct {
	Host         string `json:"host" toml:"host" yaml:"host"`
	Port         int    `json:"port" toml:"port" yaml:"port"`
	Domain       string `json:"domain" toml:"domain" yaml:"domain"`
	Path         string `json:"path" toml:"path" yaml:"path"`
	Repo         string `json:"repo" toml:"repo" yaml:"repo"`
	MaxSize      string `json:"maxSize" toml:"maxSize" yaml:"maxSize"`
	SlackChannel string `json:"slackChannel" toml:"slackChannel" yaml:"slackChannel"`
	Secret       string `json:"secret" toml:"secret" yaml:"secret"`
}

// Uploader uploader
type Uploader struct {
	c       *UploaderConfig
	s       *http.Server
	maxSize int64
}

// NewUploader create new instance for Uploader
func NewUploader(c UploaderConfig) (*Uploader, error) {
	u := &Uploader{
		c: &c,
	}
	var err error
	u.maxSize, err = ParseFileSize(c.MaxSize)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Serve set up server for upload http request
func (u *Uploader) Serve(c chan error) {
	addr := fmt.Sprintf("%s:%d", u.c.Host, u.c.Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		c <- err
		return
	}

	u.s = &http.Server{}
	mux := http.NewServeMux()
	mux.HandleFunc(u.c.Path, u.upload)
	mux.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir(u.c.Repo))))
	u.s.Handler = loggerMiddleware(mux)

	go func() {
		err = u.s.Serve(ln)
		if err != nil {
			c <- err
		}
	}()
}

// Shutdown gracefully close connections and then shutdown
func (u *Uploader) Shutdown() error {
	if u.s != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		return u.s.Shutdown(ctx)
	}
	return nil
}

func (u *Uploader) uploadNotify(filename string) error {
	if u.c.SlackChannel != "" {
		content := fmt.Sprintf("file %s upload success. <http://%s:%d/download/%s|download>.",
			filename, u.c.Domain, u.c.Port, filename)
		slackbot.Send(u.c.SlackChannel, content)
	}
	return nil
}

func (u *Uploader) upload(w http.ResponseWriter, r *http.Request) {
	if u.c.Secret != r.URL.Query().Get("secret") {
		Render(w, r, []byte("permission deny\n"), http.StatusUnauthorized)
		return
	}

	r.ParseMultipartForm(int64(u.maxSize))
	f, h, err := r.FormFile("file")
	if err != nil {
		Render(w, r, []byte(fmt.Sprintf("find upload file failed, err %v\n", err)), http.StatusBadRequest)
		return
	}
	defer f.Close()

	name := path.Join(u.c.Repo, h.Filename)
	if _, err := os.Stat(name); err != nil {
		if !os.IsNotExist(err) {
			Render(w, r, []byte(fmt.Sprintf("check file '%s' failed, err %v\n", h.Filename, err)), http.StatusBadRequest)
			return
		}
	} else {
		Render(w, r, []byte(fmt.Sprintf("file '%s' already exists\n", h.Filename)), http.StatusBadRequest)
		return
	}

	d, err := os.Create(name)
	if err != nil {
		Render(w, r, []byte(fmt.Sprintf("create file '%s' failed, err %v\n", h.Filename, err)), http.StatusInternalServerError)
		return
	}
	defer d.Close()

	n, err := io.Copy(d, f)
	if err != nil {
		os.Remove(name)
		Render(w, r, []byte(fmt.Sprintf("copy file '%s' failed, err %v\n", h.Filename, err)), http.StatusInternalServerError)
		return
	}

	if n != h.Size {
		os.Remove(name)
		Render(w, r, []byte(fmt.Sprintf("copy file '%s' not complete\n", h.Filename)), http.StatusInternalServerError)
		return
	}

	err = u.uploadNotify(h.Filename)
	if err != nil {
		Render(w, r, []byte(fmt.Sprintf("notify file '%s' upload failed, err %v\n", h.Filename, err)), http.StatusInternalServerError)
		return
	}

	Render(w, r, []byte(fmt.Sprintf("upload file '%s' success\n", h.Filename)), http.StatusOK)
	return
}
