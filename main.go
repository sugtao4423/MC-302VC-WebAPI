package main

import (
	"flag"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sugtao4423/MC-302VC-WebAPI/log"
	"github.com/sugtao4423/MC-302VC-WebAPI/mc302vc"
	"github.com/sugtao4423/MC-302VC-WebAPI/webapi"
)

func main() {
	mc302vcAddr := flag.String("addr", "", "IP address of MC-302VC")
	webApiPort := flag.String("port", "8080", "Port number of Web API")
	flag.Parse()

	if *mc302vcAddr == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Info("Starting ECHONET-Lite client...")
	mc302vc, err := mc302vc.New(*mc302vcAddr)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer mc302vc.Close()

	log.Info("Starting Web API server...")
	a := webapi.New(mc302vc)
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = a.ErrorHandler

	api := e.Group("/api")
	{
		api.GET("/status", a.GetStatus)
		api.POST("/bathAutoTimer", a.SetBathAutoTimer)
		api.POST("/bathAutoTimer/time", a.SetBathAutoTimerTime)
		api.POST("/bath/auto", a.SetBathAuto)
		api.POST("/bath/additionalHeating", a.SetBathAdditionalHeating)
	}

	err = e.Start(":" + (*webApiPort))
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
