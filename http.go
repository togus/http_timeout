package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func startHTTPServer(port, timeout int) {
	http.HandleFunc("/test", getTimeoutHandler(timeout))
	log.Println("Starting http server on port", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getTimeoutHandler(timeout int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request", r.URL, r.Host, r.Method)
		ctx := r.Context()
		keys, ok := r.URL.Query()["time"]
		var t = timeout // default timeout
		if ok {
			t, _ = strconv.Atoi(keys[0])
		}
		select {
		case <-time.After(time.Duration(t) * time.Second):
			fmt.Fprintf(w, "Waited %d seconds\n", t)
		case <-ctx.Done():
			fmt.Fprintf(w, "Error, timeout not reached")
		}
	}
}
