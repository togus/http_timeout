package main

import (
	"flag"
)

var port int
var timeout int

func main() {
	flag.IntVar(&port, "port", 8080, "Listening http port")
	flag.IntVar(&timeout, "timeout", 2, "Default timeout in seconds")
	flag.Parse()

	startHTTPServer(port, timeout)
}
