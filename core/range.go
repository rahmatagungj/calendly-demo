package core

import (
    "fmt"
    "math"
    "time"
)

func NewRange(start, end string) (Range, error) {

    t := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
    timeTemplate := "2000-01-01T%s:00.000Z"

    startTimeStr := fmt.Sprintf(timeTemplate, start)
    endTimeStr := fmt.Sprintf(timeTemplate, end)

    startTime, err := time.Parse(time.RFC3339 , startTimeStr)
    if err != nil {
        return Range{}, err
    }

    endTime, err := time.Parse(time.RFC3339, endTimeStr)
    if err != nil {
        return Range{}, err
    }

    if endTime.Before(startTime) {
        endTime = endTime.Add(24 * time.Hour)

        if endTime.Hour() > 0 {
            return Range{}, fmt.Errorf("end time must be within the same day")
        }
    }

    return Range{
        StartSec: int(math.Round(startTime.Sub(t).Seconds())),
        EndSec:   int(math.Round(endTime.Sub(t).Seconds())),
    }, nil
}

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