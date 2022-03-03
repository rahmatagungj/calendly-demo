package core

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// TODO handle use case 23 - 0

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
	// key is timestamp millis of the 00:00:00 for the given day
	DateOverrides map[int64][]Range

	// Bookings stores all booking created for this event
	Bookings Bookings

	// MaxInvitees shows maximum number of booking can be created
	MaxInvitees int
}

type GetSpotParameters struct {
	Start, End time.Time
}

func (p GetSpotParameters) IsValid() error {
	if ok := p.Start.Before(p.End); !ok {
		return fmt.Errorf("invalid date. start time must be before end")
	}
	return nil
}

type Spot struct {
	InviteeRemaining int
	StartTime        time.Time
}

func (e Event) GetAvailableSpots(params GetSpotParameters) ([]Spot, error) {

	if err := params.IsValid(); err != nil {
		return nil, err
	}

	start := params.Start.In(e.Location)
	end := params.End.In(e.Location)

	startDay := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	endDay := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())

	var spots []Spot

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
				remainingSpot := e.MaxInvitees - e.Bookings.GetBookedCount(slot)
				if remainingSpot > 0 &&
					(slot.Equal(start) || slot.After(start) && slot.Before(end)) {
					spots = append(spots, Spot{
						InviteeRemaining: remainingSpot,
						StartTime:        slot,
					})
				}
			}
		}
		curr = curr.Add(24 * time.Hour)
		if curr.After(endDay) {
			break
		}
	}
	return spots, nil
}

type CreateBookingParameters struct {
	Invitee   Invitee
	StartTime time.Time
}

var ErrTimeNotAvailable = fmt.Errorf("no time available")

// CreateBooking create new booking for given schedule if it is available
func (e *Event) CreateBooking(params CreateBookingParameters) (*Booking, error) {
	availableSpots, err := e.GetAvailableSpots(GetSpotParameters{
		Start: params.StartTime,
		End:   params.StartTime.Add(e.Duration),
	})
	if err != nil {
		return nil, err
	}

	for _, spot := range availableSpots {
		if spot.StartTime.Equal(params.StartTime) {
			b := NewBooking(params)
			e.Bookings = append(e.Bookings, *b)
			return b, nil
		}
	}
	return nil, ErrTimeNotAvailable
}
