package main

import (
	//	"encoding/json"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/jasonlvhit/gocron"
	"log"
	"os"
	//"strconv"
)

func main() {
	s := gocron.NewScheduler()
	s.Every(1).Hours().Do(CheckBalance)
	<-s.Start()
}

func CheckBalance() {
	fmt.Println("Checking dashback wallet balance.")

	// Connect to local dash RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:9998",
		User:         "dash",
		Pass:         os.Getenv("DASH_PASS"),
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current wallet balance.
	WalletBalance, err := client.GetBalance("3")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wallet balance: %d", WalletBalance)

	if WalletBalance < 5 {
		webhookUrl := "https://hooks.slack.com/services/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"

		//	attachment1 := slack.Attachment{}
		//	attachment1.AddField(slack.Field{Title: "Balance", Value: Balance})
		payload := slack.Payload{
			Text:      "Dashback wallet balance is below threshold!",
			Username:  "dashback-notifier",
			Channel:   "#dashback-notify",
			IconEmoji: ":warning:",
			//		Attachments: []slack.Attachment{attachment1},
		}
		err := slack.Send(webhookUrl, "", payload)
		if len(err) > 0 {
			fmt.Printf("error: %s\n", err)
		}
	}
}

/*
func SlackAlert() {
	webhookUrl := "https://hooks.slack.com/services/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"

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Balance", Value: "0.00000000"})
	payload := slack.Payload{
		Text:        "Dashback wallet balance is below threshold!",
		Username:    "dashback-notifier",
		Channel:     "#dashback-notify",
		IconEmoji:   ":warning:",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
*/
