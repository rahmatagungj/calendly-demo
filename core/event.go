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
    Schedules map[time.Weekday][]Range
}

type SlotParameters struct {
    month time.Month
    year  int
}

type AvailableDay struct {
    Date  time.Time
    Slots []Range
}

func (e Event) GetAvailableSlots(params SlotParameters) ([]AvailableDay, error) {
    startDayOfTheMonth := time.Date(params.year, params.month, 1, 0, 0, 0, 0, time.UTC).Day()
    endDayOfTheMonth := time.Date(params.year, params.month+1, 0, 0, 0, 0, 0, time.UTC).Day()
    var availableDays []AvailableDay
    for i := startDayOfTheMonth; i <= endDayOfTheMonth; i++ {
        now := time.Date(params.year, params.month, i, 0, 0, 0, 0, time.UTC)
        if rs, ok := e.Schedules[now.Weekday()]; ok {
            availableDays = append(availableDays, AvailableDay{
                Date:  now,
                Slots: rs,
            })
        }
    }
    return availableDays, nil
}
