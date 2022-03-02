package core_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	. "github.com/imrenagi/calendly-demo/core"
)

func TestEvent_GetAvailableSlots(t *testing.T) {

	jktTime, _ := time.LoadLocation("Asia/Jakarta")

	type fields struct {
		Event *Event
	}
	type args struct {
		params *GetSpotParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    *args
		want    []Spot
		wantErr bool
	}{
		{
			name: "should get multiple available time within user time range parameter",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 25200,
								EndSec:   28800,
							},
						},
					},
					Location: time.UTC,
					MaxInvitees: 1,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "should get multiple available time within user time range parameter",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 25200,
								EndSec:   28800,
							},
							{
								StartSec: 28800,
								EndSec:   32400,
							},
						},
					},
					Location: time.UTC,
					MaxInvitees: 1,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 7, 8, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 8, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 8, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 8, 0, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "include first available date if it is exactly the same as the user start range",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 25200,
								EndSec:   28800,
							},
						},
					},
					Location: time.UTC,
					MaxInvitees: 1,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 7, 14, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "should exclude last available if schedule",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 25200,
								EndSec:   28800,
							},
						},
					},
					Location: time.UTC,
					MaxInvitees: 1,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.February, 28, 14, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 7, 0, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "should get 2 available time within user time range parameter",
			fields: fields{
				Event: &Event{
					Duration: 30 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 0,
								EndSec:   3600,
							},
						},
					},
					Location: time.UTC,
					MaxInvitees: 1,
				},
			},

			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 7, 0, 30, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 0, 30, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 0, 30, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 0, 30, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "has 1 overrides for a date. override add some new ranges.",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 0,
								EndSec:   3600,
							},
						},
					},
					Location: time.UTC,
					DateOverrides: map[int64][]Range{
						time.Date(2022, time.February, 8, 0, 0, 0, 0, time.UTC).Unix(): []Range{
							{
								StartSec: 3600,
								EndSec:   7200,
							},
							{
								StartSec: 14400,
								EndSec:   18000,
							},
						},
					},
					MaxInvitees: 1,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 8, 1, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 8, 4, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "has 1 overrides for a date. override change existing schedule",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 0,
								EndSec:   3600,
							},
						},
					},
					Location: time.UTC,
					DateOverrides: map[int64][]Range{
						time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC).Unix(): []Range{
							{
								StartSec: 3600,
								EndSec:   7200,
							},
							{
								StartSec: 14400,
								EndSec:   18000,
							},
						},
					},
					MaxInvitees: 1,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 1, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 7, 4, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "has 1 overrides for a date. override clear an existing schedule",
			fields: fields{
				Event: &Event{
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 0,
								EndSec:   3600,
							},
						},
					},
					Location: time.UTC,
					DateOverrides: map[int64][]Range{
						time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC).Unix():  nil,
						time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC).Unix(): []Range{},
					},
					Duration: 60 * time.Minute,
					MaxInvitees: 1,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "return error if parameter is invalid",
			fields: fields{
				Event: &Event{
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should remove booked time from available slots",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 0,
								EndSec:   3600,
							},
						},
					},
					Location: time.UTC,
					Bookings: []Booking{
						{
							ID: uuid.New(),
							Invitee: Invitee{
								Email:    "foo@bar.com",
								Name:     "Foo Bar",
								Timezone: jktTime,
							},
							StartTime: time.Date(2022, time.February, 14, 7, 0, 0, 0, jktTime),
						},
					},
					MaxInvitees: 1,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
			},
			wantErr: false,
		},
		{
			name: "show remaining spots if spots still available",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 0,
								EndSec:   3600,
							},
						},
					},
					Location: time.UTC,
					Bookings: []Booking{
						{
							ID: uuid.New(),
							Invitee: Invitee{
								Email:    "foo@bar.com",
								Name:     "Foo Bar",
								Timezone: jktTime,
							},
							StartTime: time.Date(2022, time.February, 14, 7, 0, 0, 0, jktTime),
						},
					},
					MaxInvitees: 2,
				},
			},
			args: &args{
				params: &GetSpotParameters{
					Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
					End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
				},
			},
			want: []Spot{
				{StartTime: time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC), InviteeRemaining: 2},
				{StartTime: time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC), InviteeRemaining: 1},
				{StartTime: time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC), InviteeRemaining: 2},
				{StartTime: time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC), InviteeRemaining: 2},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.fields.Event.GetAvailableSpots(*tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAvailableSpots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableSpots() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSlotParameters_IsValid(t *testing.T) {
	type fields struct {
		Start time.Time
		End   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid parameter",
			fields: fields{
				Start: time.Now(),
				End:   time.Now().Add(5 * time.Minute),
			},
			wantErr: false,
		},
		{
			name: "invalid parameter",
			fields: fields{
				Start: time.Now().Add(5 * time.Minute),
				End:   time.Now(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := GetSpotParameters{
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			if err := p.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEvent_CreateBooking(t *testing.T) {

	jktTime, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)

	type fields struct {
		Event *Event
	}
	type args struct {
		params CreateBookingParameters
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantFn            func(got *Booking)
		wantBookingLength int
		wantErr           error
	}{
		{
			name: "should create booking if time is available",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 0,
								EndSec:   3600,
							},
						},
					},
					Location: time.UTC,
					MaxInvitees: 1,
				},
			},
			args: args{
				params: CreateBookingParameters{
					Invitee: Invitee{
						Email:    "foo@bar.com",
						Name:     "Foo Bar",
						Timezone: jktTime,
					},
					StartTime: time.Date(2022, 2, 7, 7, 0, 0, 0, jktTime),
				},
			},
			wantFn: func(got *Booking) {
				assert.NotEmpty(t, got.ID)
				assert.NotZero(t, got.CreatedAt)
				assert.Equal(t, time.Date(2022, 2, 7, 7, 0, 0, 0, jktTime), got.StartTime)
				assert.Equal(t, Invitee{
					Email:    "foo@bar.com",
					Name:     "Foo Bar",
					Timezone: jktTime,
				}, got.Invitee)
			},
			wantErr:           nil,
			wantBookingLength: 1,
		},
		{
			name: "should not be able to create booking if time is not available",
			fields: fields{
				Event: &Event{
					Duration: 60 * time.Minute,
					Availability: map[time.Weekday][]Range{
						time.Monday: []Range{
							{
								StartSec: 0,
								EndSec:   3600,
							},
						},
					},
					Location: time.UTC,
				},
			},
			args: args{
				params: CreateBookingParameters{
					Invitee: Invitee{
						Email:    "foo@bar.com",
						Name:     "Foo Bar",
						Timezone: jktTime,
					},
					StartTime: time.Date(2022, 2, 7, 8, 0, 0, 0, jktTime),
				},
			},
			wantFn: func(got *Booking) {
				assert.Nil(t, got)
			},
			wantErr:           ErrTimeNotAvailable,
			wantBookingLength: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.Event.CreateBooking(tt.args.params)
			assert.True(t, errors.Is(err, tt.wantErr), "CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.wantBookingLength, len(tt.fields.Event.Bookings))
			tt.wantFn(got)
		})
	}
}
