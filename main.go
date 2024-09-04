package main

import (
	"douya/admin"
	"douya/app"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	cmd := cobra.Command{}

	adminCmd := &cobra.Command{
		Use:   "admin",
		Short: "admin server start cmd",
		Run: func(cmd *cobra.Command, args []string) {
			admin.GetServer()
		},
	}

	appCmd := &cobra.Command{
		Use:   "app",
		Short: "app server start cmd",
		Run: func(cmd *cobra.Command, args []string) {
			app.GetServer()
		},
	}

	cmd.AddCommand(adminCmd)
	cmd.AddCommand(appCmd)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
