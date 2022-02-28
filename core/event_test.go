package core

import (
    "reflect"
    "testing"
    "time"
)

func TestEvent_GetAvailableSlots(t *testing.T) {

    jktTime, _ := time.LoadLocation("Asia/Jakarta")

    type fields struct {
        Event *Event
    }
    type args struct {
        params *GetSlotParameters
    }
    tests := []struct {
        name    string
        fields  fields
        args    *args
        want    []time.Time
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
                },
            },
            args: &args{
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
                },
            },
            want: []time.Time{
                time.Date(2022, time.February, 7, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 7, 0, 0, 0, time.UTC),
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
                },
            },
            args: &args{
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
                },
            },
            want: []time.Time{
                time.Date(2022, time.February, 7, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 7, 8, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 8, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 8, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 8, 0, 0, 0, time.UTC),
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
                },
            },
            args: &args{
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.February, 7, 14, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
                },
            },
            want: []time.Time{
                time.Date(2022, time.February, 7, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 7, 0, 0, 0, time.UTC),
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
                },
            },
            args: &args{
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.February, 28, 14, 0, 0, 0, jktTime),
                },
            },
            want: []time.Time{
                time.Date(2022, time.February, 7, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 7, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 7, 0, 0, 0, time.UTC),
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
                },
            },

            args: &args{
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
                },
            },
            want: []time.Time{
                time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 7,  0, 30, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 0, 30, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 0, 30, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 0, 30, 0, 0, time.UTC),
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
                },
            },
            args: &args{
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
                },
            },
            want: []time.Time{
                time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 8, 1, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 8, 4, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
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
                },
            },
            args: &args{
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
                },
            },
            want: []time.Time{
                time.Date(2022, time.February, 7, 1, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 7, 4, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
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
                        time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC).Unix(): nil,
                        time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC).Unix(): []Range{},
                    },
                    Duration: 60 * time.Minute,
                },
            },
            args: &args{
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
                },
            },
            want: []time.Time{
                time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
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
                params: &GetSlotParameters{
                    Start: time.Date(2022, time.March, 1, 0, 0, 0, 0, jktTime),
                    End:   time.Date(2022, time.February, 1, 0, 0, 0, 0, jktTime),
                },
            },
            want: nil,
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {

            got, err := tt.fields.Event.GetAvailableSlots(*tt.args.params)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetAvailableSlots() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetAvailableSlots() got = %v, want %v", got, tt.want)
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
            name:    "valid parameter",
            fields:  fields{
                Start: time.Now(),
                End:   time.Now().Add(5 * time.Minute),
            },
            wantErr: false,
        },
        {
            name:    "invalid parameter",
            fields:  fields{
                Start: time.Now().Add(5 * time.Minute),
                End:   time.Now(),
            },
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            p := GetSlotParameters{
                Start: tt.fields.Start,
                End:   tt.fields.End,
            }
            if err := p.IsValid(); (err != nil) != tt.wantErr {
                t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}