package core

import (
	"time"

	"github.com/google/uuid"
)

func NewBooking(p CreateBookingParameters) *Booking {
	return &Booking{
		ID:        uuid.New(),
		Invitee:   p.Invitee,
		StartTime: p.StartTime,
		CreatedAt: time.Now(),
	}
}

type Bookings []Booking

func (b Bookings) IsAvailable(t time.Time) bool {
	for _, booking := range b {
		if booking.StartTime.Equal(t) {
			return false
		}
	}
	return true
}

// GetBookedCount returned the total spot has been booked for a given time
func (b Bookings) GetBookedCount(t time.Time) int {
	var count int
	for _, booking := range b {
		if booking.StartTime.Equal(t) {
			count++
		}
	}
	return count
}

type Booking struct {
	ID        uuid.UUID
	Invitee   Invitee
	StartTime time.Time
	CreatedAt time.Time
}
