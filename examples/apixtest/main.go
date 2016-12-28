package main

import (
	"log"
	"net/http"

	"github.com/dshills/apix/config"
	"github.com/dshills/apix/router"
	"github.com/dshills/apix/server"
	"github.com/dshills/apix/token"
)

const prefix = "APIX"

func main() {
	//dbConfig := config.NewDatabase(prefix)
	servConfig := config.NewServer(prefix)
	server.ConfigLog(servConfig)
	if err := token.Config(servConfig); err != nil {
		log.Fatal(err)
	}
	/*
		if err := store.Config(dbConfig); err != nil {
			log.Fatal(err)
		}
	*/
	r := router.New()
	r.NotFound = http.HandlerFunc(router.NotFoundHandler)
	server.Serve(servConfig, r)
}
