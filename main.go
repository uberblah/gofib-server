package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/uberblah/gofib-server/fib"
)

const (
	EnvAddr     = "GOFIB_ADDR"
	DefaultAddr = "127.0.0.1:8080"

	EnvPath     = "GOFIB_PATH"
	DefaultPath = "/fib/"
)

func Errorf(w http.ResponseWriter, status int, format string, args ...interface{}) {
	http.Error(w, fmt.Sprintf(format, args...), status)
}

func main() {
	addr := os.Getenv(EnvAddr)
	if addr == "" {
		addr = DefaultAddr
	}

	path := os.Getenv(EnvPath)
	if path == "" {
		path = DefaultPath
	}

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		// read the request query into a string
		nStr := r.URL.Query().Get("n")
		if nStr == "" {
			Errorf(w, http.StatusBadRequest, "Parameter 'n=<>' in the query string must specify your number")
			return
		}

		// get the user's initial integer
		n0, err := strconv.Atoi(string(nStr))
		if err != nil {
			Errorf(w, http.StatusBadRequest, "Could not parse '%s' as an integer, err='%s'", nStr, err)
			return
		}

		// find the next fibonacci number
		n1, err := fib.ExhaustiveNextFib(n0, -1)
		if err != nil {
			Errorf(w, http.StatusNotFound, "Could not find the next fibonacci number to %d, err='%s'", n0, err)
			return
		}

		// respond to the user
		_, err = fmt.Fprintln(w, n1)
		if err != nil {
			log.Printf("Could not respond to message %d from client '%s', err='%s'\n",
				n0, r.RemoteAddr, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}
