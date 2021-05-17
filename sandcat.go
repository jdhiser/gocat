package main

import (
	"flag"
	"strconv"

	"github.com/mitre/gocat/core"
)

/*
These default  values can be overridden during linking - server, group, and sleep can also be overridden
with command-line arguments at runtime.
*/
var (
	key        = "JWHQZM9Z4HQOYICDHW4OCJAXPPNHBA"
	floatDNS   = "" // DNS server to use, e.g. "8.8.8.8"
	serverName = "localhost"
	server     = "http://localhost:8888"
	paw        = ""
	group      = "red"
	c2Name     = "HTTP"
	c2Key      = ""
	listenP2P  = "false" // need to set as string to allow ldflags -X build-time variable change on server-side.
	httpProxyGateway = ""
)

func main() {
	parsedListenP2P, err := strconv.ParseBool(listenP2P)
	if err != nil {
		parsedListenP2P = false
	}
	floatDNS := flag.String("floatDNS", floatDNS, "The IP of a DNS server to use")
	serverName := flag.String("serverName", serverName, "The naked hostname of the server to use, e.g. cal.mirCat.org")
	server := flag.String("server", server, "The FQDN of the server, e.g., http://cal.mirCat.org:8888/")
	httpProxyUrl :=  flag.String("httpProxyGateway", httpProxyGateway, "URL for the HTTP proxy gateway. For environments that use proxies to reach the internet.")
	paw := flag.String("paw", paw, "Optionally specify a PAW on intialization")
	group := flag.String("group", group, "Attach a group to this agent")
	c2 := flag.String("c2", c2Name, "C2 Channel for agent")
	delay := flag.Int("delay", 0, "Delay starting this agent by n-seconds")
	verbose := flag.Bool("v", false, "Enable verbose output")
	listenP2P := flag.Bool("listenP2P", parsedListenP2P, "Enable peer-to-peer receivers")

	flag.Parse()

	c2Config := map[string]string{"c2Name": *c2, "c2Key": c2Key, "httpProxyGateway": *httpProxyUrl}
	core.Core(*floatDNS, *serverName, *server, *group, *delay, c2Config, *listenP2P, *verbose, *paw)
}
