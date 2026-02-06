# Go Round Robin Load Balancer

A lightweight HTTP reverse proxy load balancer written in Go that distributes incoming requests across multiple backend servers using a round-robin algorithm.

This project is built using only Go’s standard library and is designed as a clean educational implementation of core load balancing concepts.

---

## Overview

This load balancer listens on a local port and forwards incoming HTTP requests to a pool of backend servers. Each request is routed to the next server in sequence using round-robin selection.

It demonstrates:

- Reverse proxying
- Interface-based server abstraction
- Round-robin scheduling
- Basic request forwarding
- Minimal, dependency-free design

---

## Features

- HTTP reverse proxy using `net/http/httputil`
- Round-robin load distribution
- Pluggable server interface
- Simple and extensible architecture
- Zero third-party dependencies

---

## Architecture

### Server Interface

Each backend server implements:

```go
type Server interface {
    Address() string
    IsAlive() bool
    Serve(http.ResponseWriter, *http.Request)
}

