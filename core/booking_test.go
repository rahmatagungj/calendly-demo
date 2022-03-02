package core

import (
	"testing"
	"time"
)

func TestBookings_GetBookedCount(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		b    Bookings
		args args
		want int
	}{
		{
			name: "there is no booking for given time",
			b:    Bookings{
				{
					StartTime: time.Date(2022, 1, 1, 1, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				t: time.Date(2022, 1, 1, 2, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "there found bookings for a given time",
			b:    Bookings{
				{
					Invitee:   Invitee{
						Email:    "foo@bar.com",
						Name:     "Foo Bar",
						Timezone: time.UTC,
					},
					StartTime: time.Date(2022, 1, 1, 2, 0, 0, 0, time.UTC),
				},
				{
					Invitee:   Invitee{
						Email:    "bar@foo.com",
						Name:     "Bar Foo",
						Timezone: time.UTC,
					},
					StartTime: time.Date(2022, 1, 1, 2, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				t: time.Date(2022, 1, 1, 2, 0, 0, 0, time.UTC),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetBookedCount(tt.args.t); got != tt.want {
				t.Errorf("GetBookedCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
