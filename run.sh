#!/bin/sh

go run main.go togglHours.go workdays.go -token=`cat toggl.token` -workspace=`cat workspace.id`
