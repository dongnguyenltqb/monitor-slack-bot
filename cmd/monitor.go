package cmd

import (
	"fmt"
	"thitcho/pkg/bot"
	"thitcho/pkg/monitor"
	"time"

	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use: "monitor",
	Run: func(cmd *cobra.Command, args []string) {
		bot.SetupSlackBot()
		for {
			<-time.After(3 * time.Second)
			if message, status := monitor.Check(monitor.Get()); status {
				message = "💀 STATUS : RISK \r\n " + message
				fmt.Println(message)
				fmt.Println("====> dang gui message to slack")
				message = bot.GetSlackBot().StructMonitorStatusMessage(message)
				bot.GetSlackBot().SendMessage(message)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
