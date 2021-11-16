package main

// symmetric sliding window strategy.
// each agent keeps a buffer of its last ν + 1 distinct alters with whom he interacted in the past.
func ssw(a *Agent) []*Agent {
	if len(a.Urn) > nu+1 {
		return a.Urn[len(a.Urn)-(nu+1):]
	}

	for i := 0; i < nu+1; i++ {
		add := addAgents()
		a.Urn = append(a.Urn, add)
	}
	return a.Urn[len(a.Urn)-(nu+1):]
}
