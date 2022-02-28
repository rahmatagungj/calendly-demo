package core

import (
    "reflect"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestRange_Start(t *testing.T) {
    type fields struct {
        StartSec int
        EndSec   int
    }
    tests := []struct {
        name      string
        fields    fields
        wantStart string
        wantEnd   string
    }{
        {
            name: "01:00 - 02:00",
            fields: fields{
                StartSec: 3600,
                EndSec:   7200,
            },
            wantStart: "01:00",
            wantEnd:   "02:00",
        },

        {
            name: "18:30 - 21:00",
            fields: fields{
                StartSec: 66600,
                EndSec:   75600,
            },
            wantStart: "18:30",
            wantEnd:   "21:00",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := Range{
                StartSec: tt.fields.StartSec,
                EndSec:   tt.fields.EndSec,
            }

            assert.Equal(t, tt.wantStart, r.Start())
            assert.Equal(t, tt.wantEnd, r.End())

        })
    }
}

func TestRange_Slots(t *testing.T) {
    type fields struct {
        StartSec int
        EndSec   int
    }
    type args struct {
        date     time.Time
        duration time.Duration
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []time.Time
    }{
        {
            name:    "should only return 1 time when range can only fit 1 length of duration",
            fields:  fields{
                StartSec: 0,
                EndSec:   3600,
            },
            args:    args{
                date:     time.Date(2022, time.February, 7, 10, 0, 0, 0, time.UTC),
                duration: 60 * time.Minute,
            },
            want:    []time.Time{
                time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC),
            },
        },
        {
            name:    "should return 4 schedule when range can fit the duration multiple times",
            fields:  fields{
                StartSec: 0,
                EndSec:   7200,
            },
            args:    args{
                date:     time.Date(2022, time.February, 7, 10, 0, 0, 0, time.UTC),
                duration: 30 * time.Minute,
            },
            want:    []time.Time{
                time.Date(2022, time.February, 7, 0, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 7, 0, 30, 0, 0, time.UTC),
                time.Date(2022, time.February, 7, 1, 0, 0, 0, time.UTC),
                time.Date(2022, time.February, 7, 1, 30, 0, 0, time.UTC),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := Range{
                StartSec: tt.fields.StartSec,
                EndSec:   tt.fields.EndSec,
            }
            got := r.Slots(tt.args.date, tt.args.duration)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Slots() got = %v, want %v", got, tt.want)
            }
        })
    }
}