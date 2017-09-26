package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

func main() {
	authKeyPath := flag.String("cert", "", "Path to .p8 APNSAuthKey file (Required)")
	deviceToken := flag.String("token", "", "Push token (Required)")
	topic := flag.String("topic", "", "Topic (Required)")
	flag.Parse()

	authKey, err := token.AuthKeyFromFile(*authKeyPath)
	if err != nil {
		log.Fatal("token error:", err)
	}

	token := &token.Token{
		AuthKey: authKey,
		KeyID:   "T64N7W47U9",
		TeamID:  "264H7447N5",
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = *deviceToken
	notification.Topic = *topic
	notification.Payload = []byte(`{
			"aps" : {
				"alert" : "Hello!"
			}
		}
	`)

	client := apns2.NewTokenClient(token)
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
