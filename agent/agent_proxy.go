package agent

import (
	"fmt"
	"sync"

	"github.com/mitre/gocat/output"
	"github.com/mitre/gocat/proxy"
)

func (a *Agent) ActivateP2pReceivers() {
	a.localP2pReceivers = make(map[string]proxy.P2pReceiver)
	a.localP2pReceiverAddresses = make(map[string][]string)
	a.p2pReceiverWaitGroup = &sync.WaitGroup{}
	for receiverName, p2pReceiver := range proxy.P2pReceiverChannels {
		if err := p2pReceiver.InitializeReceiver(a.server, a.beaconContact, a.p2pReceiverWaitGroup); err != nil {
			output.VerbosePrint(fmt.Sprintf("[-] Error when initializing p2p receiver %s: %s", receiverName, err.Error()))
		} else {
			output.VerbosePrint(fmt.Sprintf("[*] Initialized p2p receiver %s", receiverName))
			a.localP2pReceivers[receiverName] = p2pReceiver
			a.p2pReceiverWaitGroup.Add(1)
			a.storeLocalP2pReceiverAddresses(receiverName, p2pReceiver)
			go p2pReceiver.RunReceiver()
		}
	}
}

func (a *Agent) TerminateP2pReceivers() {
	for receiverName, p2pReceiver := range a.localP2pReceivers {
		output.VerbosePrint(fmt.Sprintf("[*] Terminating p2p receiver %s", receiverName))
		p2pReceiver.Terminate()
	}
	a.p2pReceiverWaitGroup.Wait()
}

func (a *Agent) storeLocalP2pReceiverAddresses(receiverName string, p2pReceiver proxy.P2pReceiver) {
	for _, address := range p2pReceiver.GetReceiverAddresses() {
		if _, ok := a.localP2pReceiverAddresses[receiverName]; !ok {
			a.localP2pReceiverAddresses[receiverName] = make([]string, 0)
		}
		a.localP2pReceiverAddresses[receiverName] = append(a.localP2pReceiverAddresses[receiverName], address)
	}
}