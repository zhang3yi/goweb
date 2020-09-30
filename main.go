package main

import (
	_ "goweb/docs"
	"goweb/routers"
	"time"

	"github.com/prometheus/common/log"
)

func main() {
	log.Info(time.Now().Format("2006-01-02 15:04:05") + "strart ......")

	r := routers.InitRouter()

	r.Run(":8090")
}
