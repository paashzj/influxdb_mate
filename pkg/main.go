package main

import (
	"go.uber.org/zap"
	"influxdb_mate/pkg/influx"
	"influxdb_mate/pkg/util"
	"os"
	"os/signal"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	interrupt := make(chan os.Signal, 1)
	err := influx.Start()
	if err != nil {
		util.Logger().Error("start influx server failed ", zap.Error(err))
	} else {
		os.Exit(1)
	}
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
