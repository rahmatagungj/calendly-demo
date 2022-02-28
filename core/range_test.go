package core

import (
    "testing"

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
