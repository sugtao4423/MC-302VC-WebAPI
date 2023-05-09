package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/sugtao4423/MC-302VC-WebAPI/log"
	"github.com/sugtao4423/MC-302VC-WebAPI/mc302vc"
)

func main() {
	mc302vcAddr := flag.String("addr", "", "IP address of MC-302VC")
	webApiPort := flag.String("port", "8080", "Port number of Web API")
	flag.Parse()

	if *mc302vcAddr == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Info("Starting ECHONET-Lite server...")
	mc302vc, err := mc302vc.New(*mc302vcAddr)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer mc302vc.Close()

	log.Info("Starting Web API server...")
	http.ListenAndServe(":"+(*webApiPort), nil)
}
