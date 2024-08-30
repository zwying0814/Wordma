package cmd

import (
	"github.com/spf13/cobra"
	"wordma/log"
	"wordma/server"
)

func NewStartServerCommand(app **WordmaApp) *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:     "start",
		Aliases: []string{"start"},
		Short:   "启动服务",
		Long:    Banner,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			bootstrapApp, err := server.Bootstrap()
			if err != nil {
				log.Fatalln(err)
			}
			(*app).FiberApp = bootstrapApp
		},
	}

	return serverCmd
}
