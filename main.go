package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
)

// LogEntry is a  struct that holds log data
type LogEntry struct {
	Timestamp int64  `json:"timestamp"`
	Payload   string `json:"payload"`
}

func main() {
	// flags
	host := flag.String("host", "localhost", "host to connect to")
	port := flag.String("port", "1883", "port to use")
	topic := flag.String("topic", "test", "topic to subscribe to")
	file := flag.String("file", "log.txt", "file to log to")
	flag.Parse()

	// host
	uriStr, err := url.Parse("mqtt://" + *host + ":" + *port + "/" + *topic)
	if err != nil {
		log.Fatal("could not parse server options:", err)
	}

	// subscribe & log
	go SubscribeAndLog(uriStr, *topic, *file)
	fmt.Printf("\nLogging \"mqtt://%s:%s/%s\" to \"%s\" ...\n", *host, *port, *topic, *file)
	select {}
}
