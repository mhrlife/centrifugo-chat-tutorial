package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"os"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mysql DSN", os.Getenv("MYSQL_DSN"))

		e := echo.New()

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.GET("/ok", func(c echo.Context) error {
			return c.String(200, "ok")
		})

		e.Logger.Error(e.Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
