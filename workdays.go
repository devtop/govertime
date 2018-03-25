// Copyright 2018 Tobias Ranft. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/wlbr/feiertage"
  "time"
)

var ft feiertage.Region

func WorkdaysThisMonth() int {
  y := time.Now().AddDate(0,0,-1)
  firstDayOfMonth := time.Date(y.Year(), y.Month(), 1, y.Hour(), y.Minute(), y.Second(), y.Nanosecond(), y.Location())

  return Workdays(&firstDayOfMonth, &y)
}

func WorkdaysThisYear() int {
  y := time.Now().AddDate(0,0,-1)
  firstDayOfYear := time.Date(y.Year(), 1, 1, y.Hour(), y.Minute(), y.Second(), y.Nanosecond(), y.Location())

  return Workdays(&firstDayOfYear, &y)
}

func isFeiertag(day time.Time) bool {
  isFt := false
  for _, v := range ft.Feiertage {
    if day.After(v.AddDate(0,0,0)) && day.Before(v.AddDate(0,0,1)) {
      isFt = true
    }
  }
  return isFt
}

func isWeekend(day time.Time) bool {
  weekday := day.Weekday()
  return weekday==0 || weekday==6
}

func Workdays(from *time.Time, until *time.Time) int {
  if from == nil {
    f := time.Now()
    from = &f
  }

  if until == nil {
    u := time.Now()
    until = &u
  }

  daystowork := 0
  ft = feiertage.Hamburg(until.Year())

  for current := *from; !until.Before(current); current = current.AddDate(0,0,1) {
    if !isFeiertag(current) && !isWeekend(current){
      daystowork++
    }

  }

  return daystowork
}
