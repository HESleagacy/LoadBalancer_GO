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

<img width="1366" height="768" alt="Screenshot from 2026-02-07 04-27-51" src="https://github.com/user-attachments/assets/c87a7f54-fc18-4ec3-807d-4bddfa86fe72" />
<img width="1366" height="768" alt="Screenshot from 2026-02-07 04-25-28" src="https://github.com/user-attachments/assets/83a2f025-2f73-4462-b306-72cb6a4c0f92" />
<img width="1366" height="768" alt="Screenshot from 2026-02-07 04-25-22" src="https://github.com/user-attachments/assets/04fbb027-bc4b-42fd-b203-4b04f80d0a82" />
<img width="1366" height="768" alt="Screenshot from 2026-02-07 04-25-15" src="https://github.com/user-attachments/assets/1bb37510-b938-4ae5-9db4-5bb9d051424a" />
<img width="1366" height="768" alt="Screenshot from 2026-02-07 04-25-00" src="https://github.com/user-attachments/assets/e2d3fe11-3964-4d2c-88b0-18b5e28b81e1" />

---

```go
type Server interface {
    Address() string
    IsAlive() bool
    Serve(http.ResponseWriter, *http.Request)
}
