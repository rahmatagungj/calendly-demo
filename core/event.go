package core

import (
    "time"

    "github.com/google/uuid"
)

type Event struct {
    ID   uuid.UUID
    Name string
    // Schedules stores the information about availability for
    // each day
    Schedules Schedule
}

type Schedule struct {
    Location *time.Location
    Ranges   map[time.Weekday][]Range
}

type SlotParameters struct {
    Start, End time.Time
}

func (e Event) GetAvailableSlots(params SlotParameters) ([]time.Time, error) {

    start := params.Start.In(e.Schedules.Location)
    end := params.End.In(e.Schedules.Location)

    startDay := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
    endDay := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())

    var times []time.Time

    curr := startDay
    for {
        if rs, ok := e.Schedules.Ranges[curr.Weekday()]; ok {
            for _, r := range rs {
                availability := curr.Add(time.Duration(r.StartSec) * time.Second)
                if availability == start || availability.After(start) && availability.Before(end) {
                    times = append(times, availability)
                }
            }
        }
        curr = curr.Add(24 * time.Hour)
        if curr.After(endDay) {
            break
        }
    }

    return times, nil
}
