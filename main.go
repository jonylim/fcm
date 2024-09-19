package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	input := flag.String("i", "", "JSON file containing the message parameters")
	flag.Parse()

	if *input == "" {
		log.Fatal("input text file is required\n")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	params, err := loadParamsFromFile(*input)
	if err != nil {
		log.Fatalf("error reading file: %v\n", err)
	}
	fmt.Printf("number of registration tokens: %d\n", len(params.Tokens))

	sendToFCM(ctx, params)
}

func sendToFCM(ctx context.Context, params Message) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	msgClient, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Message client: %v\n", err)
	}

	for i, token := range params.Tokens {
		var msgNotif *messaging.Notification
		var msgAndroid *messaging.AndroidConfig
		if params.Notification != nil {
			msgNotif = (*messaging.Notification)(params.Notification)
		}
		if params.Android != nil {
			msgAndroid = &messaging.AndroidConfig{
				CollapseKey: params.Android.CollapseKey,
				Priority:    params.Android.Priority,
				Data:        params.Android.Data,
			}
		}
		msg := &messaging.Message{
			Data:         params.Data,
			Notification: msgNotif,
			Android:      msgAndroid,
			Token:        token,
		}
		response, err := msgClient.Send(ctx, msg)
		if err != nil {
			fmt.Printf("[#%d] error sending message: %v\n", i+1, err)
		} else {
			fmt.Printf("[#%d] message sent: %v\n", i+1, response)
		}
	}
}

func loadParamsFromFile(filepath string) (res Message, err error) {
	reader, err := os.Open(filepath)
	if err == nil {
		defer reader.Close()
		err = json.NewDecoder(reader).Decode(&res)
	}
	return
}
