package bot

import (
	"fmt"
	"net/http"
	"sync"
	"thitcho/pkg/monitor"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type handler struct {
}

var h handler

type slackSlashQueryParams struct {
	Token        string `json:"token" form:"token"`
	Command      string `json:"command" form:"command"`
	Text         string `json:"text" form:"text"`
	ResponseURL  string `json:"response_url" form:"response_url"`
	TriggerID    string `json:"trigger_id" form:"trigger_id"`
	UserID       string `json:"user_id" form:"user_id"`
	Username     string `json:"user_name" form:"user_name"`
	TeamID       string `json:"team_id" form:"team_id"`
	EnterpriseID string `json:"enterprise_id" form:"enterprise_id"`
	ChannelID    string `json:"channel_id" form:"channel_id"`
}

func GetSlackBotHandler() *handler {
	var x sync.Once
	x.Do(func() {
		h = handler{}
	})
	return &h
}

func (h *handler) getSystemInfo(c *gin.Context) {
	var q slackSlashQueryParams
	if err := c.Bind(&q); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(q)
	message, status := monitor.Check(monitor.Get())
	systemStatus := "✅ Normal"
	if status {
		systemStatus = "💀  Risk"
	}
	c.JSON(200, gin.H{
		"blocks": []bson.M{
			bson.M{
				"type": "section",
				"text": bson.M{
					"type": "mrkdwn",
					"text": "Hi" + "<@" + q.UserID + ">",
				},
			},
			bson.M{
				"type": "section",
				"text": bson.M{
					"type": "mrkdwn",
					"text": "Status : " + systemStatus,
				},
			},
			bson.M{
				"type": "section",
				"fields": []bson.M{
					bson.M{
						"type": "mrkdwn",
						"text": `⌚️ TIME :` + time.Now().Local().String(),
					},
				},
			},
			bson.M{
				"type": "section",
				"fields": []bson.M{
					bson.M{
						"type": "mrkdwn",
						"text": message,
					},
				},
			},
		},
	})
}

func (h *handler) SetupBotHandler(app *gin.Engine) {
	botGroup := app.Group("/bot/slack")
	{
		botGroup.POST("/sysinfo", h.getSystemInfo)
	}
}
