package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	boltdb "github.com/VolantMQ/persistence-boltdb"
	"github.com/VolantMQ/volantmq"
	"github.com/VolantMQ/volantmq/auth"
	"github.com/VolantMQ/volantmq/configuration"
	"github.com/VolantMQ/volantmq/transport"
	"go.uber.org/zap"
)

var (
	logger             *zap.Logger
	externalServer     volantmq.Server
	externalAuthKey    = "external"
	externalPort       = "8883"
	externalBoltDBFile = "./external_persist.db"
	// externalCertFile   = "server.crt"
	// externalKeyFile    = "server.key"
	externalCertFile   = ""
	externalKeyFile    = ""
	internalServer     volantmq.Server
	internalAuthKey    = "internal"
	internalPort       = "1883"
	internalBoltDBFile = "./internal_persist.db"
	internalUsername   = "testusername"
	internalPassword   = "testpassword"
)

func init() {
	logger = configuration.GetLogger().Named("server")
}

func main() {
	registerExternalAuth()
	registerInternalAuth()
	go startExternalServer()
	go startInternalServer()
	watch()
	closeExternalServer()
	closeInternalServer()
}

func registerExternalAuth() {
	err := auth.Register(externalAuthKey, new(externalAuth))
	if err != nil {
		logger.Fatal("Couldn't register *external* auth provider", zap.Error(err))
	}
}

func startExternalServer() {
	serverConfig := volantmq.NewServerConfig()
	serverConfig.OfflineQoS0 = true
	serverConfig.AllowDuplicates = true
	serverConfig.Authenticators = externalAuthKey

	serverConfig.TransportStatus = func(id string, status string) {
		logger.Info("Listener status", zap.String("id", id), zap.String("status", status))
	}
	serverConfig.OnDuplicate = func(id string, allowReplace bool) {
		logger.Info("On duplicate", zap.String("id", id), zap.Bool("allowRelace", allowReplace))
	}

	persistence, err := boltdb.New(&boltdb.Config{
		File: externalBoltDBFile,
	})
	if err != nil {
		logger.Fatal("Couldn't init BoltDB persistence", zap.Error(err))
	}
	serverConfig.Persistence = persistence

	externalServer, err = volantmq.NewServer(serverConfig)
	if err != nil {
		logger.Fatal("Couldn't create server", zap.Error(err))
	}

	authMng, err := auth.NewManager(externalAuthKey)
	if err != nil {
		logger.Fatal("Couldn't register *external* auth provider", zap.Error(err))
	}

	config := transport.NewConfigTCP(
		&transport.Config{
			Port:        externalPort,
			AuthManager: authMng,
		})
	config.CertFile = externalCertFile
	config.KeyFile = externalKeyFile

	err = externalServer.ListenAndServe(config)
	if err != nil {
		logger.Fatal("Couldn't start listener", zap.Error(err))
	}
}

func closeExternalServer() {
	err := externalServer.Close()
	if err != nil {
		logger.Error("Couldn't shutdown external server", zap.Error(err))
	} else {
		logger.Info("Closed external server")
	}
}

func registerInternalAuth() {
	err := auth.Register(internalAuthKey, new(internalAuth))
	if err != nil {
		logger.Fatal("Couldn't register *internal* auth provider", zap.Error(err))
	}
}

func startInternalServer() {
	serverConfig := volantmq.NewServerConfig()
	serverConfig.OfflineQoS0 = true
	serverConfig.AllowDuplicates = true
	serverConfig.Authenticators = internalAuthKey

	serverConfig.TransportStatus = func(id string, status string) {
		logger.Info("Listener status", zap.String("id", id), zap.String("status", status))
	}
	serverConfig.OnDuplicate = func(id string, allowReplace bool) {
		logger.Info("On duplicate", zap.String("id", id), zap.Bool("allowRelace", allowReplace))
	}

	persistence, err := boltdb.New(&boltdb.Config{
		File: internalBoltDBFile,
	})
	if err != nil {
		logger.Fatal("Couldn't init BoltDB persistence", zap.Error(err))
	}
	serverConfig.Persistence = persistence

	internalServer, err = volantmq.NewServer(serverConfig)
	if err != nil {
		logger.Fatal("Couldn't create server", zap.Error(err))
	}

	authMng, err := auth.NewManager(internalAuthKey)
	if err != nil {
		logger.Fatal("Couldn't register *internal* auth provider", zap.Error(err))
	}

	config := transport.NewConfigTCP(
		&transport.Config{
			Port:        internalPort,
			AuthManager: authMng,
		})

	err = internalServer.ListenAndServe(config)
	if err != nil {
		logger.Fatal("Couldn't start listener", zap.Error(err))
	}
}

func closeInternalServer() {
	err := internalServer.Close()
	if err != nil {
		logger.Error("Couldn't shutdown internal server", zap.Error(err))
	} else {
		logger.Info("Closed internal server")
	}
}

// watch
func watch() {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGSEGV,
		syscall.SIGTERM,
		syscall.SIGSTOP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
	)

Exit:
	for {
		s := <-c
		logger.Info("Receive signal", zap.String("signal", s.String()))
		switch s {
		case syscall.SIGUSR1, syscall.SIGUSR2:
			// Do noting
		default:
			if len(c) == 0 {
				break Exit
			}
		}
	}
}

// externalAuth
type externalAuth struct {
}

func (e *externalAuth) Password(username string, password string) auth.Status {
	return auth.StatusAllow
}

func (e *externalAuth) ACL(id string, username string, topic string, accessType auth.AccessType) auth.Status {
	// Can't publish any messages
	if accessType == auth.AccessTypeWrite {
		return auth.StatusDeny
	}

	// User MessageBox -> u/${UID}
	// Public Messages -> p/xxx
	// Group Messages  -> g/${GID}

	ts := strings.Split(topic, "/")
	if len(ts) < 2 {
		return auth.StatusDeny
	}

	switch ts[0] {
	case "u":
		userID := e.UserIDByClientID(id)
		if userID != "" && userID == ts[1] {
			return auth.StatusAllow
		}
	case "p":
		return auth.StatusAllow
	case "g":
		userID := e.UserIDByClientID(id)
		if e.IsUserInGroup(userID, ts[1]) {
			return auth.StatusAllow
		}
	}
	return auth.StatusDeny
}

func (e *externalAuth) UserIDByClientID(clientID string) string {
	return clientID
}

func (e *externalAuth) IsUserInGroup(clientID, groupID string) bool {
	return true
}

// internalAuth
type internalAuth struct {
}

func (e *internalAuth) Password(username string, password string) auth.Status {
	if username == internalUsername && password == internalPassword {
		return auth.StatusAllow
	}
	return auth.StatusDeny
}

func (e *internalAuth) ACL(id string, username string, topic string, accessType auth.AccessType) auth.Status {
	return auth.StatusAllow
}
