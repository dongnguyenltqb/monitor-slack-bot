package cmd

import (
	"fmt"
	"thitcho/pkg/bot"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A server to handle slack command",
	Run: func(cmd *cobra.Command, args []string) {
		bot.SetupSlackBot()
		boot()
	},
}

func boot() {
	app := gin.Default()
	bot.GetSlackBotHandler().SetupBotHandler(app)
	port := viper.GetViper().GetString(`host.port`)
	app.Run(fmt.Sprintf(":%v", port))
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
