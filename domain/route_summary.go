package domain

import "time"

type RouteSummary struct {
	StepsCount  int
	ElapsedTime time.Duration
}
