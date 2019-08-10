package main

import (
	"baby-run/conf"
	"baby-run/run"
	"flag"
	"log"
)

//client:%d   times:%d   url:%s
var infoTemplate = "Baby-run:a simple benchmark tool.\n" +
	"client quantities:%d\n" +
	"preset times:%d seconds\n" +
	"target url:%s\n"

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		flag.Usage()
		log.Fatal("initialize failed:", err)
	}
	var config = conf.Config

	run.Start(config)
	//duration, err := run.Get(config.Url)
	//if err != nil {
	//	log.Fatal("request error:", err)
	//}
	//fmt.Printf("test complete,use duration:%vms\n", duration.Nanoseconds()/1e6)
	//fmt.Printf(infoTemplate, config.Client, config.Times, config.Url)
}
