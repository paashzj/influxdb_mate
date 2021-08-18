#!/bin/bash

nohup influxd run >$INFLUX_HOME/influx.log 2>$INFLUX_HOME/influx_error.log &