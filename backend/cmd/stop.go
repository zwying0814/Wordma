package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"wordma/log"
)

func NewStopServerCommand(app **WordmaApp) *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:     "stop",
		Aliases: []string{"stop"},
		Short:   "停止服务",
		Long:    Banner,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if (*app).FiberApp != nil {
				err := (*app).FiberApp.Shutdown()
				if err != nil {
					log.Fatal("关闭后端服务失败：", err)
				}
			} else {
				fmt.Println("后端服务未启动")
				return
			}
			fmt.Println("后端服务已关闭")
			(*app).FiberApp = nil
		},
	}

	return serverCmd
}
