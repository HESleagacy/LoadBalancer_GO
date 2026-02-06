package main

import (
	"fmt"               // Package used @lines
	"net/http"          // Package used @lines 14
	"net/http/httputil" // Package used @lines 18
	"net/url"           // Package used @lines 22
	"os"                // Package used @lines
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
