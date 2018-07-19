package main

import (
	"net/http"
	"os"
	"util/log"
)

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	port := getPort()

	log.I("start server on port:", port)
	e := http.ListenAndServe(":"+port, nil)
	if e != nil {
		log.E("failed to start server: ", e)
	}
}

func getPort() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}

	return "80"
}
