package core

import (
    "reflect"
    "testing"
    "time"

    "github.com/google/uuid"
)

func TestEvent_GetAvailableSlots(t *testing.T) {
    type fields struct {
        ID        uuid.UUID
        Name      string
        Schedules map[time.Weekday][]Range
    }
    type args struct {
        params SlotParameters
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []AvailableDay
        wantErr bool
    }{
        {
            name:    "set 1 time range in monday, should get multiples available slot",
            fields:  fields{
                Schedules: map[time.Weekday][]Range{
                    time.Monday: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
            },
            args:    args{
                params: SlotParameters{
                    month: time.February,
                    year:  2022,
                },
            },
            want:    []AvailableDay{
                {
                    Date:  time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC),
                    Slots: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
                {
                    Date:  time.Date(2022, time.February, 14, 0, 0, 0, 0, time.UTC),
                    Slots: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
                {
                    Date:  time.Date(2022, time.February, 21, 0, 0, 0, 0, time.UTC),
                    Slots: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
                {
                    Date:  time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
                    Slots: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
            },
            wantErr: false,
        },
        {
            name:    "set 1 time range in monday, should get multiples available slot (edge cases)",
            fields:  fields{
                Schedules: map[time.Weekday][]Range{
                    time.Monday: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
            },
            args:    args{
                params: SlotParameters{
                    month: time.December,
                    year:  2022,
                },
            },
            want:    []AvailableDay{
                {
                    Date:  time.Date(2022, time.December, 5, 0, 0, 0, 0, time.UTC),
                    Slots: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
                {
                    Date:  time.Date(2022, time.December, 12, 0, 0, 0, 0, time.UTC),
                    Slots: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
                {
                    Date:  time.Date(2022, time.December, 19, 0, 0, 0, 0, time.UTC),
                    Slots: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
                {
                    Date:  time.Date(2022, time.December, 26, 0, 0, 0, 0, time.UTC),
                    Slots: []Range{
                        {
                            StartSec: 25200,
                            EndSec:   28800,
                        },
                    },
                },
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
            }
            got, err := e.GetAvailableSlots(tt.args.params)
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
