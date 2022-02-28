package core

import (
    "time"
)

type Range struct {
    StartSec, EndSec int
}

func (r Range) Start() string {
    return r.intToString(r.StartSec)
}

func (r Range) End() string {
    return r.intToString(r.EndSec)
}

func (r Range) intToString(s int) string {
    t := time.Unix(int64(s), 0).In(time.UTC)
    return t.Format("15:04")
}

// Slots return all start time that is available for the range [startTime, endTime)
func (r Range) Slots(date time.Time, duration time.Duration) []time.Time {
    cur := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
    start := cur.Add(time.Duration(r.StartSec) * time.Second)
    end := cur.Add(time.Duration(r.EndSec) * time.Second)

    var availabilities []time.Time
    curr := start
    for {
        if curr == end || curr.After(end) {
            break
        }
        end := curr.Add(duration)
        if end == end || end.Before(end) {
            availabilities = append(availabilities, curr)
        }
        curr = curr.Add(duration)
    }
    return availabilities
}