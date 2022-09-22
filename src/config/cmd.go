package config

import "github.com/spf13/cobra"

var initCommand = cobra.Command{
	Short: "simple api with golang",
}

func init() {
	initCommand.AddCommand(ServeCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
