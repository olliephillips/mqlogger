package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// SubscribeAndLog creates a subscription and logs received
// messages to a file with received timestamp
func SubscribeAndLog(uri *url.URL, topic string, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("there was a problem opening file:", err)
	}

	client := connect("mqlogger", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		m := new(LogEntry)
		m.Timestamp = time.Now().Unix()
		m.Payload = string(msg.Payload())
		output, err := json.Marshal(m)
		if err != nil {
			log.Println("problem marshaling to json")
		}

		output = append(output, "\n"...)
		if _, err := f.Write(output); err != nil {
			log.Fatal("problem writing to file:", err)
		}
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func connect(clientID string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientID, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	return opts
}
