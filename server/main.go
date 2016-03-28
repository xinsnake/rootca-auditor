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
	listen := flag.Bool("listen", true, "whether to listen to a local port")

	flag.Parse()

	if *refresh {
		//TODO other platforms
		ProcessOSX()
	}

	fs := http.FileServer(http.Dir("wwwroot"))
	http.Handle("/", fs)

	if *listen {
		log.Printf("Listening on port %d ...\n", *port)
		http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	}
}
