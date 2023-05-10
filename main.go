package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sugtao4423/MC-302VC-WebAPI/log"
	"github.com/sugtao4423/MC-302VC-WebAPI/mc302vc"
	"github.com/sugtao4423/MC-302VC-WebAPI/public"
	"github.com/sugtao4423/MC-302VC-WebAPI/webapi"
)

func main() {
	mc302vcAddr := flag.String("addr", "", "IP address of MC-302VC")
	webApiPort := flag.String("port", "8080", "Port number of Web API")
	webUser := flag.String("user", "", "Username of Web API")
	webPass := flag.String("pass", "", "Password of Web API")
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
	a := webapi.New(mc302vc, *webUser, *webPass)
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = a.ErrorHandler
	e.Validator = a.NewJsonValidator()

	fs := http.FileServer(public.Root)
	e.GET("/*", echo.WrapHandler(fs))

	api := e.Group("/api")
	{
		if *webUser != "" && *webPass != "" {
			api.Use(a.BasicAuthMiddleware())
		}
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
