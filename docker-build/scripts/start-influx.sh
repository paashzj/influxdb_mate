#!/bin/bash

mkdir -p $INFLUX_HOME/logs
nohup influxd run >>$INFLUX_HOME/logs/influx.stdout.log 2>>$INFLUX_HOME/logs/influx.stderr.log &

