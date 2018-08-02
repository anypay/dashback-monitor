package main

import "github.com/ashwanthkumar/slack-go-webhook"
import "fmt"

func main() {
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
