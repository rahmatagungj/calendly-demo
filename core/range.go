package core

import (
    "time"
)

type Range struct {
    StartSec, EndSec int
}

func (r Range) Start() string {
    return r.intToString(r.StartSec)
}

func (r Range) End() string {
    return r.intToString(r.EndSec)
}

func (r Range) intToString(s int) string {
    t := time.Unix(int64(s), 0).In(time.UTC)
    return t.Format("15:04")
}
