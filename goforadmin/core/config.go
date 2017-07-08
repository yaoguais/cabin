package core

import (
	"fmt"
	"github.com/Yaoguais/gadmin/lib/db"
	"github.com/Yaoguais/gadmin/lib/log"
	"github.com/Yaoguais/gadmin/lib/random"
	"github.com/Yaoguais/gadmin/lib/redis"
	"github.com/spf13/viper"
	"qiniupkg.com/api.v7/conf"
	"strings"
)

var (
	configFile    string
	appName       string
	appStaticHost string
	appHost       string
	appRoot       string
	logPath       string
	logLevel      int
	dbConf        *db.DbConf
	redisConf     *redis.RedisConf
	sessionDriver string
	sessionPath   string
	sessionSecret string
	uploadBucket  string
	uploadDomain  string
	coreView      string
)

func initConfig() {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		panic("read config file failed " + err.Error())
	}
	// [app]
	if appName = viper.GetString("app.name"); len(appName) == 0 {
		panic("app name can't be empty")
	}
	if appStaticHost = viper.GetString("app.static_host"); len(appStaticHost) == 0 {
		panic("app static host can't be empty")
	}
	if appHost = viper.GetString("app.host"); len(appHost) == 0 {
		panic("app host can't be empty")
	}
	// [log]
	if logPath = viper.GetString("log.path"); len(logPath) == 0 {
		panic("log path can't be empty")
	}
	logLevel = log.Ldebug
	logLvlName := strings.ToUpper(viper.GetString("log.level"))
	if logLvlName == "DEBUG" {
		logLevel = log.Ldebug
	} else if logLvlName == "INFO" {
		logLevel = log.Linfo
	} else if logLvlName == "WARN" {
		logLevel = log.Lwarn
	} else if logLvlName == "ERROR" {
		logLevel = log.Lerror
	} else if logLvlName == "PANIC" {
		logLevel = log.Lpanic
	} else if logLvlName == "FATAL" {
		logLevel = log.Lfatal
	} else {
		panic("log level should be DEBUG/INFO/WARN/ERROR/PANIC/FATAL")
	}
	// [db]
	var dbDriver, dbHost, dbUser, dbPassword, dbDatabase, dbCharset string
	var dbPort int
	if dbDriver = viper.GetString("db.driver"); len(dbDriver) == 0 {
		panic("db driver can't be empty")
	}
	if dbHost = viper.GetString("db.host"); len(dbHost) == 0 {
		panic("db host can't be empty")
	}
	if dbPort = viper.GetInt("db.port"); dbPort <= 0 {
		panic("db port should be greater than 0")
	}
	if dbUser = viper.GetString("db.user"); len(dbUser) == 0 {
		panic("db user can't be empty")
	}
	dbPassword = viper.GetString("db.password")
	if dbDatabase = viper.GetString("db.database"); len(dbDatabase) == 0 {
		panic("db database can't be empty")
	}
	if dbCharset = viper.GetString("db.charset"); len(dbCharset) == 0 {
		panic("db driver can't be empty")
	}

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase,
		dbCharset,
	)
	dbConf = &db.DbConf{
		Driver: dbDriver,
		DSN:    dbDSN,
	}
	// [session]
	sessionDriver = viper.GetString("session.driver")
	if sessionPath = viper.GetString("session.path"); len(sessionPath) == 0 {
		panic("session path can't be empty")
	}
	if sessionSecret = viper.GetString("session.secret"); len(sessionSecret) == 0 {
		sessionSecret = random.RandowmString(18)
	}
	// [qiniu]
	if qiniuAccessKey := viper.GetString("qiniu.access_key"); len(qiniuAccessKey) != 0 {
		if qiniuSecretKey := viper.GetString("qiniu.secret_key"); len(qiniuSecretKey) != 0 {
			conf.ACCESS_KEY = qiniuAccessKey
			conf.SECRET_KEY = qiniuSecretKey
		}
	}
	uploadBucket = viper.GetString("qiniu.bucket")
	uploadDomain = viper.GetString("qiniu.domain")
	// [core]
	coreView = viper.GetString("core.view")
	// [redis]
	var (
		redisHost, redisPassword string
		redisPort, redisDatabase int
	)
	if redisHost = viper.GetString("redis.host"); len(redisHost) == 0 {
		panic("redis host can't be empty")
	}
	if redisPort = viper.GetInt("redis.port"); redisPort <= 0 {
		panic("redis port should be greater than 0")
	}
	redisPassword = viper.GetString("redis.password")
	if redisDatabase = viper.GetInt("redis.database"); redisDatabase < 0 {
		panic("redis port should be greater than -1")
	}
	redisConf = &redis.RedisConf{
		Address:  fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword,
		Database: redisDatabase,
	}
}
