#!/bin/sh

go run main.go hoursYesterdaysMonth.go -token=`cat toggl.token` -workspace=`cat workspace.id`
