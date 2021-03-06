package main

import (
    "fmt"
    "time"

    "github.com/imrenagi/calendly-demo/core"
)

func main() {
    e := core.Event{
        Location: time.UTC,
        Availability: map[time.Weekday][]core.Range{
            time.Sunday: []core.Range{{StartSec: 3600, EndSec: 7200}, {StartSec: 7200, EndSec: 10800}},
            time.Monday: []core.Range{{StartSec: 3600, EndSec: 7200}},
            // time.Tuesday: []core.Range{{StartSec: 3600, EndSec: 7200}},
            // time.Wednesday: []core.Range{{StartSec: 3600, EndSec: 7200}},
            // time.Thursday: []core.Range{{StartSec: 3600, EndSec: 7200}},
            // time.Friday: []core.Range{{StartSec: 3600, EndSec: 7200}},
            // time.Saturday: []core.Range{{StartSec: 3600, EndSec: 7200}},
        },
        Duration: 30 * time.Minute,
    }

    loc, _ := time.LoadLocation("Asia/Jakarta")
    ts, _ := e.GetAvailableSpots(core.GetSpotParameters{
        Start: time.Date(2022, 2, 1, 8, 0, 0, 0, loc),
        End: time.Date(2022, 2, 28, 8, 0, 0, 0, loc),
    })

    for _, t := range ts {
        fmt.Println(t)
    }
}
