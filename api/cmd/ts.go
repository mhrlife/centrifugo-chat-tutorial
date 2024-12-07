package cmd

import (
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/serializer"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"

	"github.com/spf13/cobra"
)

var types = []any{
	serializer.UserInfo{},
	serializer.UserWithToken{},
	serializer.RegisterRequest{},
}

var tsCmd = &cobra.Command{
	Use:   "ts",
	Short: "Export type script types",
	Run: func(cmd *cobra.Command, args []string) {
		converter := typescriptify.New().
			WithInterface(true)

		for _, a := range types {
			converter = converter.Add(a)
		}

		err := converter.ConvertToFile("../front/src/types/serializer.ts")
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(tsCmd)
}
