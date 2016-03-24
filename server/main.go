package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := flag.Int("port", 8999, "port the static server will listen on")
	refresh := flag.Bool("refresh", true, "refresh local certificate data")

	flag.Parse()

	if *refresh {
		ProcessOSX()
		//TODO other platforms
	}

	fs := http.FileServer(http.Dir("wwwroot"))
	http.Handle("/", fs)

	log.Printf("Listening on port %d ...\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
