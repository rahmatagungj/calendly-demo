package core

import (
    "reflect"
    "testing"
    "time"

    "github.com/google/uuid"
)

func TestEvent_GetAvailableSlots(t *testing.T) {

    jktTime, _ := time.LoadLocation("Asia/Jakarta")

    type fields struct {
        ID        uuid.UUID
        Name      string
        Schedules Schedule
        Duration  time.Duration
        DateOverrides map[int64][]Range
    }
    type args struct {
        params *SlotParameters
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
                Schedules:
                Schedule{
                    Ranges: map[time.Weekday][]Range{
                        time.Monday: []Range{
                            {
                                StartSec: 25200,
                                EndSec:   28800,
                            },
                        },
                    },
                    Location: time.UTC,
                },
                Duration: 60 * time.Minute,
            },
            args: &args{
                params: &SlotParameters{
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
                Schedules:
                Schedule{
                    Ranges: map[time.Weekday][]Range{
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
                Duration: 60 * time.Minute,
            },
            args: &args{
                params: &SlotParameters{
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
                Schedules:
                Schedule{
                    Ranges: map[time.Weekday][]Range{
                        time.Monday: []Range{
                            {
                                StartSec: 25200,
                                EndSec:   28800,
                            },
                        },
                    },
                    Location: time.UTC,
                },
                Duration: 60 * time.Minute,
            },
            args: &args{
                params: &SlotParameters{
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
                Schedules:
                Schedule{
                    Ranges: map[time.Weekday][]Range{
                        time.Monday: []Range{
                            {
                                StartSec: 25200,
                                EndSec:   28800,
                            },
                        },
                    },
                    Location: time.UTC,
                },
                Duration: 60 * time.Minute,
            },
            args: &args{
                params: &SlotParameters{
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
                Schedules:
                Schedule{
                    Ranges: map[time.Weekday][]Range{
                        time.Monday: []Range{
                            {
                                StartSec: 0,
                                EndSec:   3600,
                            },
                        },
                    },
                    Location: time.UTC,
                },
                Duration: 30 * time.Minute,
            },
            args: &args{
                params: &SlotParameters{
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
                Schedules: Schedule{
                    Ranges: map[time.Weekday][]Range{
                        time.Monday: []Range{
                            {
                                StartSec: 0,
                                EndSec:   3600,
                            },
                        },
                    },
                    Location: time.UTC,
                },
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
                Duration: 60 * time.Minute,
            },
            args: &args{
                params: &SlotParameters{
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
                Schedules: Schedule{
                    Ranges: map[time.Weekday][]Range{
                        time.Monday: []Range{
                            {
                                StartSec: 0,
                                EndSec:   3600,
                            },
                        },
                    },
                    Location: time.UTC,
                },
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
                Duration: 60 * time.Minute,
            },
            args: &args{
                params: &SlotParameters{
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
                Schedules: Schedule{
                    Ranges: map[time.Weekday][]Range{
                        time.Monday: []Range{
                            {
                                StartSec: 0,
                                EndSec:   3600,
                            },
                        },
                    },
                    Location: time.UTC,
                },
                DateOverrides: map[int64][]Range{
                    time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC).Unix(): nil,
                    time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC).Unix(): []Range{},
                },
                Duration: 60 * time.Minute,
            },
            args: &args{
                params: &SlotParameters{
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
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {

            e := Event{
                ID:        tt.fields.ID,
                Name:      tt.fields.Name,
                Schedules: tt.fields.Schedules,
                Duration: tt.fields.Duration,
                DateOverrides: tt.fields.DateOverrides,
            }
            got, err := e.GetAvailableSlots(*tt.args.params)
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
