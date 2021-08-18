package influx

import (
	"influxdb_mate_go/pkg/path"
	"influxdb_mate_go/pkg/util"
)

func Start() error {
	return util.CallScript(path.InfluxStartScript)
}
