package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/uberblah/gofib-server/fib"
)

const EnvAddr = "GOFIB_ADDR"

func Errorf(w http.ResponseWriter, status int, format string, args ...interface{}) {
	http.Error(w, fmt.Sprintf(format, args...), status)
}

func main() {
	addr := os.Getenv(EnvAddr)
	if addr == "" {
		log.Fatalf("Env variable %s must be set, like '127.0.0.1:8080'", EnvAddr)
	}

	http.HandleFunc("/fib/", func(w http.ResponseWriter, r *http.Request) {

		// read the request into a string
		bodyStr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Errorf(w, http.StatusBadRequest, "Failed to read the body of your request, err='%s'")
			return
		}

		// get the user's initial integer
		n0, err := strconv.Atoi(string(bodyStr))
		if err != nil {
			Errorf(w, http.StatusBadRequest, "Could not parse '%s' as an integer, err='%s'", bodyStr, err)
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
