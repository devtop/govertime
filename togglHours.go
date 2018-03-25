// Copyright 2018 Tobias Ranft. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/devtop/go-toggl-reports/togglreports"
	"os"
	"time"
)

type TogglReportConnection struct {
	workspaceID int
	client *togglreports.Client
}

func (c *TogglReportConnection) Init(apiToken string, workspaceID int) *TogglReportConnection {
	c.workspaceID = workspaceID
	c.client = togglreports.NewClient(apiToken)

	return c
}

func (c *TogglReportConnection) HoursThisMonth() *time.Duration {
	y := time.Now().Add(time.Hour * 24 * -1)
	start := time.Date(y.Year(), y.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(y.Year(), y.Month(), y.Day(), 23, 59, 59, 0, time.UTC)

	return c.getFromUntil(&start, &end)
}

func (c *TogglReportConnection) HoursThisYear() *time.Duration {
	y := time.Now().Add(time.Hour * 24 * -1)
	start := time.Date(y.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(y.Year(), y.Month(), y.Day(), 23, 59, 59, 0, time.UTC)

	return c.getFromUntil(&start, &end)
}

func (c *TogglReportConnection) getDuration(selection *togglreports.Selectparameters) time.Duration {

	r, err := c.client.Summary.Get(c.workspaceID, selection)
	checkError(err)

	return time.Duration(r.TotalGrand) * time.Millisecond
}

func (c *TogglReportConnection) getFromUntil(start *time.Time, end *time.Time) *time.Duration {
	selection := &togglreports.Selectparameters{
		Start:       start,
		End:         end,
	}

	d := c.getDuration(selection)

  selection = &togglreports.Selectparameters{
		Start:       start,
		End:         end,
    Description:  "Fahrtzeit",
	}

  fz := c.getDuration(selection)

  total := d - ( fz / 2 )

  return &total
}




func TogglHours(apiToken string, workspaceID int, selection *togglreports.Selectparameters) (time.Duration){

	c := togglreports.NewClient(apiToken)

	r, err := c.Summary.Get(workspaceID, selection)
	checkError(err)

	return time.Duration(r.TotalGrand) * time.Millisecond
}

func TogglHoursFromUntil(apiToken string, workspaceID int, start *time.Time, end *time.Time) *time.Duration {

	selection := &togglreports.Selectparameters{
		Start:       start,
		End:         end,
	}

	d := TogglHours(apiToken, workspaceID, selection)

  selection = &togglreports.Selectparameters{
		Start:       start,
		End:         end,
    Description:  "Fahrtzeit",
	}

  fz := TogglHours(apiToken, workspaceID, selection)

  total := d - ( fz / 2 )

  return &total
}

func TogglHoursThisMonth(apiToken string, workspaceID int)  (*time.Duration){

	y := time.Now().Add(time.Hour * 24 * -1)
	start := time.Date(y.Year(), y.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(y.Year(), y.Month(), y.Day(), 23, 59, 59, 0, time.UTC)

	return TogglHoursFromUntil(apiToken, workspaceID, &start, &end)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
