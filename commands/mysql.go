package commands

import (
	"github.com/orderbynull/one/core"
	"github.com/orderbynull/one/providers"
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Use MySQL as lock provider",
	Run: func(cmd *cobra.Command, args []string) {
		provider := providers.NewMySQLLock()
		core.Process(provider, "xxx", "ls -la; sleep 3")
	},
}

func init() {
	RootCmd.AddCommand(mysqlCmd)
}
