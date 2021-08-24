package influx

import (
	"influxdb_mate/pkg/path"
	"influxdb_mate/pkg/util"
	"time"
)

func Start() error {
	err := util.CallScript(path.InfluxStartScript)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Second)
	err = util.CallScript(path.InfluxSetupScript)
	InitClient()
	return err
}
