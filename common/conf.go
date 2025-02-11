package common

import (
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
)

var ConfHandle Conf
var Log slog.Logger

type SvrConf struct {
	Address string
	Port    int
}

type LoggingConf struct {
	Level string
}

type PersistConf struct {
	Driver        string
	SvrAddr       string
	SvrPort       string
	SvrApiKey     string
	SvrUserName   string
	SvrPassword   string
	StorageBucket string
	OrgName       string
}

type DougaInfoSaveTriggerConf struct {
	Enabled                bool
	MongoSvrConnURI        string
	DbName                 string
	ACIDInfoCollectionName string
}

type PollerConf struct {
	Interval  uint
	RandShift bool
}

type DougaInfoSvrConf struct {
	Addr     string
	Port     string
	UserName string
	Password string
}

type RPCConf struct {
	DougaInfo DougaInfoSvrConf
}

type Conf struct {
	Server        SvrConf
	Logging       LoggingConf
	Poller        PollerConf
	Persist       PersistConf
	RPC           RPCConf
	DougaInfoSave DougaInfoSaveTriggerConf
}

func logLevelMap(str string) slog.Level {
	var logMap map[string]slog.Level = map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}
	return logMap[str]
}

func InitLogger() {
	Log = *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevelMap(ConfHandle.Logging.Level),
	}))
	Log.Info("Init Logger Level: " + ConfHandle.Logging.Level)
}

func InitConfDriver(confPath string) {
	if confPath == "" {
		confPath = "./conf.toml"
	}
	if _, err := toml.DecodeFile(confPath, &ConfHandle); err != nil {
		panic(err)
	}
}
