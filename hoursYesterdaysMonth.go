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

func HoursYesterdaysMonth(apiToken string, workspaceID int)  (*time.Duration){

	c := togglreports.NewClient(apiToken)


	y := time.Now().Add(time.Hour * 24 * -1)
	start := time.Date(y.Year(), y.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(y.Year(), y.Month(), y.Day(), 23, 59, 59, 0, time.UTC)

	selection := &togglreports.Selectparameters{
		Start:       &start,
		End:         &end,
	}

	s, err := c.Summary.Get(workspaceID, selection)
	checkError(err)

	d := time.Duration(s.TotalGrand) * time.Millisecond

  selection = &togglreports.Selectparameters{
		Start:       &start,
		End:         &end,
    Description:  "Fahrtzeit",
	}

	f, err := c.Summary.Get(workspaceID, selection)
	checkError(err)

  fz := time.Duration(f.TotalGrand) * time.Millisecond

  total := d - ( fz / 2 )

  return &total
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
