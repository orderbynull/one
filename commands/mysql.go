package commands

import (
	"fmt"
	"github.com/orderbynull/one/core"
	"github.com/orderbynull/one/providers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

const bindTmpl = "mysql.%s"
const mysqlCmdUse = "mysql [flags] [command]"
const mysqlCmdShort = "Use MySQL as lock provider"
const mysqlCmdFlagUser = "user"
const mysqlCmdFlagHost = "host"
const mysqlCmdFlagPort = "port"
const mysqlCmdFlagPassword = "password"

// mysqlCmdFlags holds MySQL specific flags.
var mysqlCmdFlags = []string{mysqlCmdFlagUser, mysqlCmdFlagPassword, mysqlCmdFlagHost, mysqlCmdFlagPort}

// mysqlCmd is cobra-command for executing user command wrapped by MySQL lock.
var mysqlCmd = &cobra.Command{
	Use:   mysqlCmdUse,
	Short: mysqlCmdShort,
	Args:  cobra.MinimumNArgs(1),
	Run:   mysqlCmdRun,
}

// mysqlCmdRun is Run-function for mysqlCmd cobra-command.
var mysqlCmdRun = func(cmd *cobra.Command, args []string) {
	command := strings.Join(args, " ")

	name, err := core.MakeLockName(command)
	if err != nil {
		core.ErrorAndExit(err)
	}

	if err := core.Process(initMySQLLocker(), name, command); err != nil {
		core.ErrorAndExit(err)
	}
}

// initMySQLLocker returns *MySQLLock filled with dynamic parameters.
// Each parameter value can arrive either from config or from command flag.
var initMySQLLocker = func() *providers.MySQLLock {
	return providers.NewMySQLLock(
		viper.GetString(fmt.Sprintf(bindTmpl, mysqlCmdFlagUser)),
		viper.GetString(fmt.Sprintf(bindTmpl, mysqlCmdFlagPassword)),
		viper.GetString(fmt.Sprintf(bindTmpl, mysqlCmdFlagHost)),
		viper.GetString(fmt.Sprintf(bindTmpl, mysqlCmdFlagPort)),
	)
}

// init defines available flags and binds them to Viper config.
func init() {
	for _, flag := range mysqlCmdFlags {
		mysqlCmd.Flags().String(flag, "", "")
		viper.BindPFlag(fmt.Sprintf(bindTmpl, flag), mysqlCmd.Flags().Lookup(flag))
	}

	RootCmd.AddCommand(mysqlCmd)
}
