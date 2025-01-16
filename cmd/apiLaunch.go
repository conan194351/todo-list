package cmd

import (
	"github.com/conan194351/todo-list.git/internal/api"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var apiLaunchCmd = &cobra.Command{
	Use:   "api:launch",
	Short: "Launch API server",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			api.ServerModule,
			fx.Invoke(func(api *api.Server) {
				api.Run()
			}),
		).Run()
	},
}

func init() {
	rootCmd.AddCommand(apiLaunchCmd)
}
