package main

import (
	"fmt"
	"math/rand"
)

type Agent struct {
	Id  int
	Urn []*Agent
}

// for debug
func (a *Agent) String() string {
	str := fmt.Sprintf("{ Id: %d, Urn: [", a.Id)
	for _, c := range a.Urn {
		str += fmt.Sprintf("%d, ", c.Id)
	}
	str += "]}"
	return str
}

// core of agents behavior. the from-agent and to-agent will interact with strategy s
func (from *Agent) interact(s func(*Agent) []*Agent) (*Agent, *Agent) {
	// pick opponent agent from to-agent's urn
	i := rand.Intn(len(from.Urn))
	to := from.Urn[i]

	// memory buffer
	fromMb := make([]*Agent, 0)
	toMb := make([]*Agent, 0)

	// cleaning memory buffer
	// if memory buffer has the opponent, remove it
	for _, a := range s(from) {
		if a.Id != to.Id {
			fromMb = append(fromMb, a)
		}
	}
	for _, a := range s(to) {
		if a.Id != from.Id {
			toMb = append(toMb, a)
		}
	}

	// add ρ to-agent to the from-agent's urn, vise-versa
	for i := 0; i < rho; i++ {
		from.Urn = append(from.Urn, to)
		to.Urn = append(to.Urn, from)
	}

	// exchange memory buffer
	from.Urn = append(from.Urn, toMb...)
	to.Urn = append(to.Urn, fromMb...)

	return from, to
}
