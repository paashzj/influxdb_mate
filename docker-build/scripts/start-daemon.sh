#!/bin/bash


nohup $INFLUX_HOME/mate/influxdb_mate >$INFLUX_HOME/influxdb_mate.stdout.log 2>$INFLUX_HOME/influxdb_mate.stderr.log
