package agent

// Creates and initializes a new Agent. Upon success, returns a pointer to the agent and nil Error.
// Upon failure, returns nil and an error.
func AgentFactory(floatDNS, serverName, server string, group string, c2Config map[string]string, enableLocalP2pReceivers bool, initialDelay int, paw string) (*Agent, error) {
	newAgent := &Agent{}
	if err := newAgent.Initialize(floatDNS, serverName, server, group, c2Config, enableLocalP2pReceivers, initialDelay, paw); err != nil {
		return nil, err
	} else {
		newAgent.Sleep(newAgent.initialDelay)
		return newAgent, nil
	}
}
