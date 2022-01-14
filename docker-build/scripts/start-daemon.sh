#!/bin/bash

mkdir -p $INFLUX_HOME/logs
nohup $INFLUX_HOME/mate/influxdb_mate >>$INFLUX_HOME/logs/influxdb_mate.stdout.log 2>>$INFLUX_HOME/logs/influxdb_mate.stderr.log

