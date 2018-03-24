#!/bin/sh

go run main.go hoursYesterdaysMonth.go workdays.go -token=`cat toggl.token` -workspace=`cat workspace.id`
