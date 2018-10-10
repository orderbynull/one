package commands

import "github.com/spf13/cobra"

var name string

var RootCmd = &cobra.Command{
	Use:   "one",
	Short: "Short description",
	Long:  "Longer description feel free to use a few lines here.",
}

func init() {
	RootCmd.PersistentFlags().StringVar(&name, "name", "", "lock name")
}
