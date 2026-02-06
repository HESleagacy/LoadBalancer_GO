package main

import (
	"fmt"               // Package used @lines 33
	"net/http"          // Package used @lines 14, 54
	"net/http/httputil" // Package used @lines 18
	"net/url"           // Package used @lines 22
	"os"                // Package used @lines 34
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request) // forwarded request
}
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy // Pointer to reverse proxy instance
}

func newSimpleServer(addr string) *simpleServer { // Create simpleServer
	serverUrl, err := url.Parse(addr)
	handleErr(err) // func @ line 31

	return &simpleServer{ // Struct defined @line 16
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

type LoadBalancer struct {
	port            string
	roundRobinCount int      // Counter
	servers         []Server // List of Servers
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{ // Struct defined @line 38
		port:            port,
		roundRobinCount: 0, // Initially 0
		servers:         servers,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server {}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {} // Proxy Handler Method
