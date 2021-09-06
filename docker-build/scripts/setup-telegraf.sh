#!/bin/bash

influx org create -n telegraf
influx bucket create -n telegraf -o telegraf -r 72h
bucketId=`influx -o telegraf bucket list|grep telegraf|awk '{print $1}'`
influx auth create -o telegraf --read-bucket $bucketId --write-bucket $bucketId