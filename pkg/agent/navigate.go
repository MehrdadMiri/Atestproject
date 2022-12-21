package agent

import (
	"Atestproject/pkg/command"
)

func (a *Agent) Navigate(l command.Location) bool { // Returns true if agent is already in location
	if a.Location.Equals(l) {
		return true
	}

	return false
}
