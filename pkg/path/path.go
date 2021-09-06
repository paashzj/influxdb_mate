package path

import (
	"os"
	"path/filepath"
)

// influx
var (
	InfluxHome = os.Getenv("INFLUX_HOME")
)

// mate
var (
	InfluxMatePath                 = filepath.FromSlash(InfluxHome + "/mate")
	InfluxScripts                  = filepath.FromSlash(InfluxMatePath + "/scripts")
	InfluxStartScript              = filepath.FromSlash(InfluxScripts + "/start-influx.sh")
	InfluxSetupScript              = filepath.FromSlash(InfluxScripts + "/setup-influx.sh")
	InfluxSetupTelegrafTokenScript = filepath.FromSlash(InfluxScripts + "/setup-telegraf.sh")
)
