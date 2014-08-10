package fitters

import (
	"errors"
	"time"
)

type window struct {
	weights    []uint
	cycle_time time.Duration
	resolution time.Duration
	buckets    uint
}

type WindowFit interface {
	Pack(Request) bool
	Buckets() uint
}

func (w *window) Buckets() uint {
	return w.buckets
}

func (w *window) Pack(r Request) bool {
	return true
}

func NewWindow(cycle_time time.Duration, resolution time.Duration) (w WindowFit, err error) {
	if cycle_time < resolution {
		return nil, errors.New("Can't specify cycle_time < resolution")
	}

	wp := &window{cycle_time: cycle_time, resolution: resolution, buckets: (uint)(cycle_time / resolution)}

	wp.weights = make([]uint, wp.buckets)
	return wp, nil
}
