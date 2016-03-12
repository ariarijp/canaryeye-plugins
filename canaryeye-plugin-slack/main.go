package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/ariarijp/canaryeye"
	"github.com/bluele/slack"
)

func main() {
	if len(os.Getenv("CANARYEYE_PLUGIN_SLACK_TOKEN")) == 0 {
		err := errors.New("CANARYEYE_PLUGIN_SLACK_TOKEN env var must be set")
		canaryeye.HandleError(err)
	}

	if len(os.Getenv("CANARYEYE_PLUGIN_SLACK_CHANNEL")) == 0 {
		err := errors.New("CANARYEYE_PLUGIN_SLACK_CHANNEL env var must be set")
		canaryeye.HandleError(err)
	}

	token := os.Getenv("CANARYEYE_PLUGIN_SLACK_TOKEN")
	channelName := os.Getenv("CANARYEYE_PLUGIN_SLACK_CHANNEL")

	r := canaryeye.GetResultSlice(os.Stdin)
	var mBuf bytes.Buffer

	mBuf.WriteString("```\r\n")
	for _, result := range r.Results {
		mBuf.WriteString(fmt.Sprintf("%s\t%d\r\n", result.Host, result.Count))
	}
	mBuf.WriteString("```")

	api := slack.New(token)

	channel, err := api.FindChannelByName(channelName)
	canaryeye.HandleError(err)

	err = api.ChatPostMessage(channel.Id, mBuf.String(), nil)
	canaryeye.HandleError(err)
}
