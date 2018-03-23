# gofib-server
A simple HTTP server which, given one fibonacci number, produces the next

## Prerequisites
- latest go version
- `${GOPATH}/bin` must be in your `$PATH`

## Installation
The following will install gofib-server in `${GOPATH}/bin`
```bash
go install github.com/uberblah/gofib-server
```

## Running
The following will serve over HTTP on `localhost:8080/fib/`
```bash
gofib-server
```

## Manual testing
```bash
# should print 3
curl localhost:8080/fib/?n=2

# should print 5
curl localhost:8080/fib/?n=3

# should print an error indicating that 4 is not a fibonacci number
curl localhost:8080/fib/?n=4
```

## Overriding the local resource path and hostname
```bash
# starts the server, reachable at 192.168.0.13:8080/fib/
GOFIB_ADDR=192.168.0.13:8080 gofib-server

# starts the server, reachable at localhost:8080/fibonoid/
GOFIB_PATH=/fibonoid/ gofib-server

# starts the server, reachable at 192.168.0.13:8080/fibonoid/
GOFIB_ADDR=192.168.0.13:8080 GOFIB_PATH=/fibonoid/ gofib-server
```
