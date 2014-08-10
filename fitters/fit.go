package fitters

import (
	"time"
)

type Request struct {
	Workload uint
	// A rate function of uints of work Per time.Duration
	Units uint
	Per   time.Duration
}
