package main

import "flag"

var port int

func main() {
	flag.IntVar(&port, "post", 8080, "Listening http port")

	flag.Parse()

}