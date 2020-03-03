package main

import (
	"fmt"

	"github.com/atreya2011/slack"
)

func main() {
	api := slack.New("YOUR_TOKEN_HERE")

	userID := "USER_ID"

	_, _, channelID, err := api.OpenIMChannel(userID)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	api.PostMessage(channelID, slack.MsgOptionText("Hello World!", false))
}
