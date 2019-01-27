package db

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/tarantool/go-tarantool"

	"github.com/mikhailbadin/csp-aggregator/config"
)

type dbStore struct {
	mongo     *mgo.Session
	tarantool *tarantool.Connection
}

const (
	MongoDBName         = "csp"
	MongoCollectionName = "reports"
)

const (
	SpaceScriptSrc = "script_src"
)

var db *dbStore

// Init is initialize MongoDB and Tarantool
func Init() error {
	mongo, err := mongoInit()
	if err != nil {
		return fmt.Errorf("cannot initialize MongoDB: %s", err.Error())
	}
	tarantool, err := tarantoolInit()
	if err != nil {
		return fmt.Errorf("cannot initialize Tarantool: %s", err.Error())
	}
	db = &dbStore{
		mongo:     mongo,
		tarantool: tarantool,
	}
	return nil
}

// GetMongoDB get MongoDB connection
func GetMongoDB() *mgo.Session {
	return db.mongo
}

// GetTarantoolDB get Tarantool connection
func GetTarantoolDB() *tarantool.Connection {
	return db.tarantool
}

func mongoInit() (*mgo.Session, error) {
	conf := config.GetMongoConf()
	conn, err := mgo.Dial(conf.URL)
	if err != nil {
		return nil, fmt.Errorf("cannot connect: %s", err.Error())
	}
	return conn, nil
}

func tarantoolInit() (*tarantool.Connection, error) {
	conf := config.GetTarantoolConf()
	opts := tarantool.Opts{
		Timeout:       conf.Timeout,
		Reconnect:     conf.Reconnect,
		MaxReconnects: conf.MaxReconnects,
		User:          conf.User,
		Pass:          conf.Pass,
	}
	conn, err := tarantool.Connect(conf.URL, opts)
	if err != nil {
		return nil, fmt.Errorf("cannot connect: %s", err.Error())
	}
	return conn, nil
}
