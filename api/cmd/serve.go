package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"os"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mysql DSN", os.Getenv("MYSQL_DSN"))

		fmt.Println(echo.New().Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
