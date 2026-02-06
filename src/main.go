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

func newSimpleServer(addr string) *simpleServer { // Create Server
	serverUrl, err := url.Parse(addr)
	handleErr(err) // func @ line 31

	return &simpleServer{ // Struct defined @line 16
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

// Methods
func (s *simpleServer) Address() string { return s.addr }

func (s *simpleServer) IsAlive() bool { return true }

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) { //forward user req
	s.proxy.ServeHTTP(rw, req)
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

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)] // Cycled-indexing

	for !server.IsAlive() { // until any liver server is found
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}

	lb.roundRobinCount++
	return server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{ // Struct defined @line 38
		port:            port,
		roundRobinCount: 0, // Initially 0
		servers:         servers,
	}
}
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()                              // Gets the next sever @ line 53
	fmt.Printf("forwarding request to address %q\n", targetServer.Address()) // add error handling later
	targetServer.Serve(rw, req)
}

func main() {
	servers := []Server{ // Example servers
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("http://www.bing.com"),
		newSimpleServer("http://www.duckduckgo.com"),
	}

	lb := NewLoadBalancer("3000", servers) // lb @port 8000

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
