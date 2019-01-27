package server

import "github.com/mikhailbadin/csp-aggregator/config"

// Init is initialize server
func Init() {
	conf := config.GetServerConf()
	router := NewRouter()
	router.Run(conf.Addr)
}
