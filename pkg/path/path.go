package path

import (
	"os"
	"path/filepath"
)

// zookeeper
var (
	InfluxHome           = os.Getenv("INFLUX_HOME")
)

// mate
var (
	InfluxMatePath              = filepath.FromSlash(InfluxHome + "/mate")
	InfluxScripts               = filepath.FromSlash(InfluxMatePath + "/scripts")
	InfluxStartScript           = filepath.FromSlash(InfluxScripts + "/start-influx.sh")
)
