package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

const listCmdUse = "list"
const listCmdTmpl = "> %s\n"
const listCmdHead = "Available lock providers:"
const listCmdShort = "List available lock providers"

// listCmdProviders holds list of possible lock providers.
var listCmdProviders = []string{"mysql"}

// listCmd cobra-command prints list of possible lock providers.
var listCmd = &cobra.Command{
	Use:   listCmdUse,
	Short: listCmdShort,
	Run:   listCmdRun,
}

// listCmdRun is Run-function for listCmd cobra-command
var listCmdRun = func(cmd *cobra.Command, args []string) {
	println(listCmdHead)
	for _, provider := range listCmdProviders {
		fmt.Printf(listCmdTmpl, provider)
	}
}

func init() {
	RootCmd.AddCommand(listCmd)
}
