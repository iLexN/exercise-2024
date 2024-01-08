/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	//	"go1/services/logger"
	"go1/services/route_service"
)

// serverStartCmd represents the serverStart command
var serverStartCmd = &cobra.Command{
	Use:   "server:start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverStart()
	},
}

func init() {
	rootCmd.AddCommand(serverStartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverStartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverStartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serverStart() {
	router := gin.Default()

	// set trusted proxies , nil to  disable
	trustedProxies := []string{"127.0.0.1"}
	err := router.SetTrustedProxies(trustedProxies)
	if err != nil {
		fmt.Print(err)
		return
	}

	route_service.CreateRoute(router)

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
