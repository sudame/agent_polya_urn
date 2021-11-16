package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/cheggaaa/pb/v3"
	_ "github.com/mattn/go-sqlite3"
)

var events = make([]Event, 0)

var agents = make([]*Agent, 0)

// drip active (= has more than 0 agents in the urn) agents
// from the environment
func activeAgents() []*Agent {
	ret := make([]*Agent, 0)
	for _, a := range agents {
		if len(a.Urn) > 0 {
			ret = append(ret, a)
		}
	}
	return ret
}

// add agents to environment
func addAgents() *Agent {
	a := Agent{
		Id: len(agents),
	}
	agents = append(agents, &a)
	return &a
}

// setup initial 2 agents
func setupAgents() {
	// create initial 2 agents
	a := addAgents()
	b := addAgents()

	// add to urn each other
	a.Urn = append(a.Urn, b)
	b.Urn = append(b.Urn, a)
}

func pickRandomActiveAgents() *Agent {
	u := make([]*Agent, 0)
	for _, a := range activeAgents() {
		for i := 0; i < len(a.Urn); i++ {
			u = append(u, a)
		}
	}

	i := rand.Intn(len(u))

	return u[i]
}

// read CLI arguments and set global vars
func parseArgs() {
	var (
		_rho  = flag.Int("rho", 1, "ρ")
		_nu   = flag.Int("nu", 1, "ν")
		_iter = flag.Int("iter", 1000, "iteration")
	)

	flag.Parse()
	rho = *_rho
	nu = *_nu
	iter = *_iter
}

func main() {
	// prepare
	parseArgs()
	setupAgents()

	// main experiment
	bar := pb.StartNew(iter)
	for i := 0; i < iter; i++ {
		from := pickRandomActiveAgents()
		_, to := from.interact(ssw)
		event := Event{From: from.Id, To: to.Id}
		events = append(events, event)

		bar.Increment()
	}
	bar.Finish()

	// dump result
	os.Mkdir("result", os.ModePerm)
	dumpEventLog(fmt.Sprintf("result/event__rho%d_nu%d_iter%d.dat", rho, nu, iter))
	dumpAgents(fmt.Sprintf("result/agents__rho%d_nu%d_iter%d.dat", rho, nu, iter))
	aggEdges(fmt.Sprintf("result/edges__rho%d_nu%d_iter%d.csv", rho, nu, iter))
	aggNodes(fmt.Sprintf("result/nodes__rho%d_nu%d_iter%d.csv", rho, nu, iter))
}
