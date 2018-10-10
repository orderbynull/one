package commands

import (
	"github.com/orderbynull/one/core"
	"github.com/orderbynull/one/providers"
	"github.com/spf13/cobra"
	"strings"
)

const nameFlag = "name"
const usage = "mysql [flags] [command]"
const description = "Use MySQL as lock provider"

var mysqlCmd = &cobra.Command{
	Use:   usage,
	Short: description,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag(nameFlag).Value.String()
		command := strings.Join(args, " ")

		if name == "" {
			name = core.MakeLockName(command)
		}

		provider := providers.NewMySQLLock("root", "515528aA", "127.0.0.1", 3306)
		core.Process(provider, name, command)
	},
}

func init() {
	RootCmd.AddCommand(mysqlCmd)
}
