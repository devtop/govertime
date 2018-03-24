// Copyright 2018 Tobias Ranft. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
//  "github.com/wlbr/feiertage"
  "time"
)

func workdays(from *time.Time, until *time.Time) int {
  if from == nil {
    f := time.Now()
    from = &f
  }

  if until == nil {
    u := time.Now()
    until = &u
  }

  daystowork := 0
  //ft := feiertage.Hamburg(until.Year())

  for current := *from; !until.Before(current); current = current.AddDate(0,0,1) {
    if true {
      daystowork++
    }

  }

  return daystowork
}
