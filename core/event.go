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

    // Duration defines how long an event should take
    Duration time.Duration

    DateOverrides map[int64][]Range
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

        if dateOverrides, ok := e.DateOverrides[curr.Unix()]; ok {
            for _, r := range dateOverrides {
                // TODO DRY this
                startAvailable := curr.Add(time.Duration(r.StartSec) * time.Second)
                endAvailable := curr.Add(time.Duration(r.EndSec) * time.Second)

                slots, err := e.slotsInRange(startAvailable, endAvailable)
                if err != nil {
                    return nil, err
                }

                for _, slot := range slots {
                    if slot == start || slot.After(start) && slot.Before(end) {
                        times = append(times, slot)
                    }
                }
            }
        } else {
            if rs, ok := e.Schedules.Ranges[curr.Weekday()]; ok {
                for _, r := range rs {
                    startAvailable := curr.Add(time.Duration(r.StartSec) * time.Second)
                    endAvailable := curr.Add(time.Duration(r.EndSec) * time.Second)

                    slots, err := e.slotsInRange(startAvailable, endAvailable)
                    if err != nil {
                        return nil, err
                    }

                    for _, slot := range slots {
                        if slot == start || slot.After(start) && slot.Before(end) {
                            times = append(times, slot)
                        }
                    }
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

// slotsInRange return all start time that is available for the range [startTime, endTime)
func (e Event) slotsInRange(startTime, endTime time.Time) ([]time.Time, error) {
    var availabilities []time.Time
    curr := startTime
    for {
        if curr == endTime || curr.After(endTime) {
            break
        }
        end := curr.Add(e.Duration)
        if end == endTime || end.Before(endTime) {
            availabilities = append(availabilities, curr)
        }
        curr = curr.Add(e.Duration)
    }
    return availabilities, nil
}
