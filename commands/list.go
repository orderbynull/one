package commands

import "github.com/spf13/cobra"

var echoCmd = &cobra.Command{
	Use:   "list",
	Short: "List available lock providers",
	Run: func(cmd *cobra.Command, args []string) {
		println("Available lock providers:")
		println("- mysql")
		println("- ports")
		println("- flock")
		println("- redis")
	},
}

func init() {
	RootCmd.AddCommand(echoCmd)
}
