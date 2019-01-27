package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Server environment definition
const (
	serverAddrEnv = "SERVER_ADDR"
)

// MongoDB environment definition
const (
	mongoURLEnv = "MONGO_URL"
)

// TarantoolDB environment definition
const (
	tarantoolURLEnv          = "TARANTOOL_URL"
	tarantoolUserEnv         = "TARANTOOL_USER"
	tarantoolPassEnv         = "TARANTOOL_PASS"
	tarantoolTimeoutEnv      = "TARANTOOL_TIMEOUT"
	tarantoolReconnectEnv    = "TARANTOOL_RECONNECT"
	tarantoolMaxReconnectEnv = "TARANTOOL_RECONNECT_MAX"
)

// Server configuration struct
type Server struct {
	Addr string
}

// MongoDB configuration struct
type MongoDB struct {
	URL string
}

// TarantoolDB configuration struct
type TarantoolDB struct {
	URL           string
	User          string
	Pass          string
	Timeout       time.Duration
	Reconnect     time.Duration
	MaxReconnects uint
}

type config struct {
	mongo     *MongoDB
	tarantool *TarantoolDB
	server    *Server
}

var conf *config

// GetServerConf return Server config
func GetServerConf() *Server {
	return conf.server
}

// GetMongoConf return MongoDB config
func GetMongoConf() *MongoDB {
	return conf.mongo
}

// GetTarantoolConf return Tarantool config
func GetTarantoolConf() *TarantoolDB {
	return conf.tarantool
}

// Init configuration
func Init() error {
	mongo, err := mongoInit()
	if err != nil {
		return fmt.Errorf("cannot init MongoDB config: %s", err.Error())
	}
	tarantool, err := tarantoolInit()
	if err != nil {
		return fmt.Errorf("cannot init Tarantool config: %s", err.Error())
	}
	server, err := serverInit()
	if err != nil {
		return fmt.Errorf("cannot init Server config: %s", err.Error())
	}
	conf = &config{
		mongo:     mongo,
		tarantool: tarantool,
		server:    server,
	}
	return nil
}

func mongoInit() (*MongoDB, error) {
	url := os.Getenv(mongoURLEnv)
	if url == "" {
		return nil, fmt.Errorf("environment '%s' not specified", mongoURLEnv)
	}
	return &MongoDB{
		URL: url,
	}, nil
}

func tarantoolInit() (*TarantoolDB, error) {
	timeout := time.Millisecond * 3000
	reconnect := time.Second * 10
	var reconnectMax uint = 10
	url := os.Getenv(tarantoolURLEnv)
	if url == "" {
		return nil, fmt.Errorf("environment '%s' not specified", tarantoolURLEnv)
	}
	user := os.Getenv(tarantoolUserEnv)
	if user == "" {
		return nil, fmt.Errorf("environment '%s' not specified", tarantoolUserEnv)
	}
	pass := os.Getenv(tarantoolPassEnv)
	timeoutStr := os.Getenv(tarantoolTimeoutEnv)
	if timeoutStr != "" {
		timeoutUint, err := strconv.ParseUint(timeoutStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("cannot parse environment '%s' as uint32", tarantoolTimeoutEnv)
		}
		timeout = time.Millisecond * time.Duration(timeoutUint)
	}
	reconnectStr := os.Getenv(tarantoolReconnectEnv)
	if reconnectStr != "" {
		reconnectUint, err := strconv.ParseUint(reconnectStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("cannot parse environment '%s' as uint32", tarantoolReconnectEnv)
		}
		reconnect = time.Second * time.Duration(reconnectUint)
	}
	reconnectMaxStr := os.Getenv(tarantoolMaxReconnectEnv)
	if reconnectMaxStr != "" {
		reconnectMaxUint, err := strconv.ParseUint(reconnectMaxStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("cannot parse environment '%s' as uint32", tarantoolMaxReconnectEnv)
		}
		reconnectMax = uint(reconnectMaxUint)
	}
	return &TarantoolDB{
		URL:           url,
		User:          user,
		Pass:          pass,
		Timeout:       timeout,
		Reconnect:     reconnect,
		MaxReconnects: reconnectMax,
	}, nil
}

func serverInit() (*Server, error) {
	addr := os.Getenv(serverAddrEnv)
	if addr == "" {
		addr = ":8080"
	}
	return &Server{
		Addr: addr,
	}, nil
}
