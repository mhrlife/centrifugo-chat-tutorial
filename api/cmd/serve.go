package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"os"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		urlPrefix := os.Getenv("URL_PREFIX")

		e := echo.New()

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.Pre(middleware.Rewrite(map[string]string{
			urlPrefix + "/*": "/$1",
			urlPrefix:        "/",
		}))

		e.GET("/ok", func(c echo.Context) error {
			return c.String(200, "ok")
		})

		e.Logger.Errorf("url prefix is %s", urlPrefix)

		e.Logger.Error(e.Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
