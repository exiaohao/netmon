package agent

import "log"

type Agent struct {
	AgentID	string
}

func (a *Agent) Run(stopCh <- chan struct{}) {
	log.Printf("Agent starting, AgentID=%s", a.AgentID)
}