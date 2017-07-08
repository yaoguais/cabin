package session

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"

	"gopkg.in/redis.v5"
)

var (
	store         sessions.Store
	sessionDriver string
	sessionPath   string
	sessionSecret string
	redisClient   *redis.Client
)

func SetSessionPath(path string) {
	sessionPath = path
}

func SetSessionSecret(secret string) {
	sessionSecret = secret
}

func SetSessionDriver(driver string) {
	sessionDriver = driver
}

func SetRedisClient(client *redis.Client) {
	redisClient = client
}

func StartSession() {

	if sessionDriver == "file" {
		fileStore := sessions.NewFilesystemStore(sessionPath, []byte(sessionSecret))
		store = fileStore
	} else if sessionDriver == "redis" {
		redisStore, _ := NewRediStore(redisClient, []byte(sessionSecret))
		store = redisStore
	} else {
		panic("session driver should be file/redis, current is " + sessionDriver)
	}
}

func SessionName() string {
	return "session_id"
}

func GetSession(context echo.Context, key string) *sessions.Session {
	session, _ := store.Get(context.Request(), key)

	return session
}

func SetSession(context echo.Context, session *sessions.Session) {
	store.Save(context.Request(), context.Response().Writer, session)
}
