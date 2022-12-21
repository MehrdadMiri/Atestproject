package agent

import (
	"Atestproject/pkg/command"
	"math"
	"time"
)

type Agent struct {
	Destinations []command.Location // Follow up destination
	command.Location
	ID              int
	WaitTime        time.Duration
	WaitingDuration time.Duration
	IsReturning     bool
	ReturnLocation  command.Location
	ETA             time.Duration
}

func New(id int, home command.Location, wait time.Duration) *Agent {
	return &Agent{
		ID:              id,
		ReturnLocation:  home,
		WaitingDuration: wait,
	}
}

func GetETA(a, b command.Location) time.Duration {
	distance := a.Distance(b)
	eta := math.Ceil(float64(distance) * float64(time.Second))
	return time.Duration(int64(eta))
}

func (a *Agent) RemainingTime() time.Duration {
	remainingTime := a.ETA
	if len(a.Destinations) > 0 {
		remainingTime += GetETA(a.Location, a.Destinations[0])
	}
	for i := 0; i < len(a.Destinations)-1; i++ {
		remainingTime += GetETA(a.Destinations[i], a.Destinations[i+1])
	}
	return remainingTime
}

func (a *Agent) LastLocation() command.Location {
	if len(a.Destinations) > 0 {
		return a.Destinations[len(a.Destinations)-1]
	}
	return a.Location
}

func (a *Agent) Cycle(t time.Duration) { // TODO refactor
	switch {
	case a.ETA > 0:
		if a.ETA < t {
			eta := a.ETA
			a.ETA = 0
			a.Cycle(t - eta) // Todo: Remove self call
		} else {
			a.ETA -= t
		}
	case a.IsReturning:
		a.Location = a.ReturnLocation
		a.ETA = GetETA(a.Location, a.ReturnLocation)
		a.IsReturning = false
		a.WaitTime = 0
		a.Cycle(t)
	case len(a.Destinations) > 0:
		a.Location = a.Destinations[0]
		a.ETA = GetETA(a.Location, a.Destinations[0])
		a.Destinations = a.Destinations[1:]
		a.WaitTime = 0
		a.Cycle(t)
	case a.WaitTime+t <= a.WaitingDuration:
		a.WaitTime += t
	case a.WaitTime+t > a.WaitingDuration:
		a.IsReturning = true
		a.WaitTime = 0
		a.Cycle(t - (a.WaitingDuration - a.WaitTime))
	}
}

func (a *Agent) AddLocation(l command.Location) {
	a.Destinations = append(a.Destinations, l)
}
