package main

import "loadbalancer/loadbalancer"

func main() {
	loadbalancer.MakeLoadBalancer(5)
	// servers.RunServers(5)
}
