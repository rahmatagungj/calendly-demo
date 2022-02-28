package core

import (
    "fmt"
    "time"

    "github.com/google/uuid"
)

type Event struct {
	ID   uuid.UUID
	Name string

	// Location defines the timezone used by calendar creator
    Location *time.Location

	// Duration defines how long an event should take
	Duration time.Duration

    // Availability stores the information about availability for
    // each day
    Availability map[time.Weekday][]Range

	// DateOverrides specify the overriding range for a specific day
	// key is timestamp milis of the 00:00:00 for the given day
	DateOverrides map[int64][]Range
}

type GetSlotParameters struct {
	Start, End time.Time
}

func (p GetSlotParameters) IsValid() error {
    if ok := p.Start.Before(p.End); !ok {
        return fmt.Errorf("invalid date. start time must be before end")
    }
    return nil
}

func (e Event) GetAvailableSlots(params GetSlotParameters) ([]time.Time, error) {

    if err := params.IsValid(); err != nil {
        return nil, err
    }

	start := params.Start.In(e.Location)
	end := params.End.In(e.Location)

	startDay := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	endDay := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())

	var times []time.Time

	curr := startDay
	for {
		var ranges []Range
		if dateOverrides, ok := e.DateOverrides[curr.Unix()]; ok {
			ranges = dateOverrides
		} else if scheduleRanges, ok := e.Availability[curr.Weekday()]; ok {
			ranges = scheduleRanges
		}

		for _, r := range ranges {
            for _, slot := range r.Slots(curr, e.Duration) {
                if slot == start || slot.After(start) && slot.Before(end) {
                    times = append(times, slot)
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
