// Copyright 2018 Tobias Ranft. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
  "flag"
)

// run with
// go run main.go hoursYesterdaysMonth.go -token=`cat toggl.token` -workspace=`cat workspace.id`
func main() {

  apiToken := flag.String("token", "foo", "Toggle API Token")
  workspaceID := flag.Int("workspace", 0, "Your Workspace ID")

  flag.Parse()

  s := HoursYesterdaysMonth(*apiToken, *workspaceID)

  fmt.Println(s.String())

	fmt.Println(workdays())

}
