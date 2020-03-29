package bot

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type SlackBot struct {
	Token           string
	MonitorChannnel string
	LogChannel      string
	PostMessageURL  string
}

var slackBot SlackBot

func SetupSlackBot() {
	fmt.Println("Setting up Slack Bot")
	slackBot = SlackBot{
		Token:           viper.GetViper().GetString(`bot.slack.token`),
		MonitorChannnel: viper.GetViper().GetString(`bot.slack.monitor_channel`),
		LogChannel:      viper.GetViper().GetString(`bot.slack.log_channel`),
		PostMessageURL:  viper.GetViper().GetString(`bot.slack.url.post_message`),
	}
	fmt.Println("Setup slack bot sucessfully", slackBot)
}

func GetSlackBot() *SlackBot {
	return &slackBot
}

func (bot *SlackBot) StructMonitorStatusMessage(message string) string {
	channel := bot.MonitorChannnel
	str := `{
  "channel": "` + channel + `",
  "blocks": [
	{
      "type": "section",
      "fields": [
        {
          "type": "mrkdwn",
          "text": "Hey <!here>"
        }
      ]
    },
	{
      "type": "section",
      "fields": [
        {
          "type": "mrkdwn",
		  "text": "⌚️ TIME : ` + time.Now().Local().String() + `"
        }
      ]
    },
    {
      "type": "section",
      "fields": [
        {
          "type": "mrkdwn",
          "text":"`
	str += message
	str += `" }
      ]
    }
  ]
}`
	return str
}

func (bot *SlackBot) SendMessage(message string) {
	token := bot.Token
	url := bot.PostMessageURL
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(message)))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}
