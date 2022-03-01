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

type Booking struct {
	ID        uuid.UUID
	Invitee   Invitee
	StartTime time.Time
	CreatedAt time.Time
}
