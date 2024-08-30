package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"wordma/log"
)

type WordmaApp struct {
	RootCmd  *cobra.Command
	FiberApp *fiber.App
}

var Version = "0.0.1(2024-08-23)"
var Banner = `
██╗    ██╗ ██████╗ ██████╗ ██████╗ ███╗   ███╗ █████╗ 
██║    ██║██╔═══██╗██╔══██╗██╔══██╗████╗ ████║██╔══██╗
██║ █╗ ██║██║   ██║██████╔╝██║  ██║██╔████╔██║███████║
██║███╗██║██║   ██║██╔══██╗██║  ██║██║╚██╔╝██║██╔══██║
╚███╔███╔╝╚██████╔╝██║  ██║██████╔╝██║ ╚═╝ ██║██║  ██║
 ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝

Wordma ` + Version + `

 -> 一个自部署的静态站点扩展工具包
 -> https://blog.zwying.com
`

func App() *WordmaApp {
	cmd := &WordmaApp{
		RootCmd: &cobra.Command{
			Use:     "wdm",
			Short:   "Wordma：一个自部署的静态站点扩展工具包",
			Long:    Banner,
			Version: Version,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(Banner)
				fmt.Print("-------------------------------\n\n")
				fmt.Println("提示: 在命令结尾添加 `-h` 标志，可查看帮助信息")
			},
		},
	}
	cmd.RootCmd.CompletionOptions.DisableDefaultCmd = true
	cmd.RootCmd.SetHelpTemplate(HelpTemplate())
	cmd.RootCmd.SetVersionTemplate("Wordma {{printf \"%s\" .Version}}\n")

	return cmd
}

func (wdm *WordmaApp) LunchApp() error {
	// 挂载命令
	wdm.addCommand(NewStartServerCommand(&wdm))
	wdm.addCommand(NewStopServerCommand(&wdm))
	wdm.addCommand(NewCreateAdminCommand())

	done := make(chan bool, 1) // shutdown signal

	// listen for interrupt signal to gracefully shutdown the application
	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		<-sigch
		done <- true
	}()

	// ===================
	//  2. Command Execute
	// ===================
	go func() {
		if err := wdm.RootCmd.Execute(); err != nil {
			color.Red(err.Error())
		}
		done <- true
	}()
	<-done

	// 关闭后端服务

	if wdm.FiberApp != nil {
		err := wdm.FiberApp.Shutdown()
		if err != nil {
			log.Fatal("关闭后端服务失败：", err)
		}
	}

	return nil
}

func (wdm *WordmaApp) addCommand(cmd *cobra.Command) {
	originalPreRunFunc := cmd.PreRun

	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		if originalPreRunFunc != nil {
			originalPreRunFunc(cmd, args) // extends original pre-run logic
		}
	}

	wdm.RootCmd.AddCommand(cmd)
}
