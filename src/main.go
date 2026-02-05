package main

import(
	"net/http/httputil"                      // Package used @lines 
	"net/url"                                // Package used @lines
)

type simpleServer struct{                    // Structure of the Server
	addr string                              // Address of the Server
	proxy "httputil.ReverseProxy"            // ReverseProxy to hide sever port from client
}