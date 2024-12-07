package cmd

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mhrlife/centrifugo-chat-tutorial/config"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/endpoint"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/ent"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.New()
		if err != nil {
			logrus.WithError(err).Fatal("failed to load config")
		}

		logrus.Info("config loaded")

		client, err := ent.Open("mysql", cfg.Database.DSN)
		if err != nil {
			logrus.WithError(err).Fatal("failed opening connection to mysql")
		}

		logrus.Info("mysql connection established")

		if err := client.Schema.Create(context.Background()); err != nil {
			logrus.WithError(err).Fatal("failed creating schema resources")
		}

		logrus.Info("schema created")

		svc := service.NewService(client, cfg)
		defer svc.Close()

		echoServer := echo.New()

		echoServer.Use(middleware.Recover())
		echoServer.Use(middleware.Logger())

		endpointManager := endpoint.NewEndpoint(cfg, echoServer, svc)
		endpointManager.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
