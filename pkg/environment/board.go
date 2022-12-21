package environment

import (
	"Atestproject/pkg/agent"
	"Atestproject/pkg/command"
	"time"
)

type Env struct {
	Agents                 []*agent.Agent
	UnassignedDestinations []command.Location
}

func New(agentsCount int, agentWait time.Duration, home command.Location) *Env {
	agents := make([]*agent.Agent, agentsCount)
	for i := range agents {
		agents[i] = agent.New(i, home, agentWait)
	}
	destinations := make([]command.Location, 0)
	return &Env{agents, destinations}
}

func (e *Env) AddNewLocation(l command.Location) {
	bestAgent := 0
	bestEta := time.Duration(1000 * 1000 * 1000) // or infinite duration
	for i, a := range e.Agents {
		agentEta := a.RemainingTime() + agent.GetETA(l, a.LastLocation())
		if agentEta < bestEta {
			bestAgent = i
			bestEta = agentEta
		}
	}
	e.Agents[bestAgent].AddLocation(l)
}

func (e *Env) Cycle(t time.Duration) {
	for _, a := range e.Agents {
		a.Cycle(t)
	}
}
