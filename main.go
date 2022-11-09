package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vipxkw/proxypool/api"
	"github.com/vipxkw/proxypool/internal/app"
	"github.com/vipxkw/proxypool/internal/cron"
	"github.com/vipxkw/proxypool/pkg/proxy"
)

var configFilePath = ""

func main() {
	flag.StringVar(&configFilePath, "c", "", "path to config file: config.yaml")
	flag.Parse()

	if configFilePath == "" {
		configFilePath = os.Getenv("CONFIG_FILE")
	}
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	err := app.InitConfigAndGetters(configFilePath)
	if err != nil {
		panic(err)
	}

	proxy.InitGeoIpDB()
	fmt.Println("Do the first crawl...")
	go app.CrawlGo()
	go cron.Cron()
	api.Run()
}
