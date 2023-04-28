package version

import (
	"fmt"
	"github.com/duktig666/go-common/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const CLI_VERSION = "0.0.1"

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: config.GlobalConfig.Cli.Name + " version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Printf(config.GlobalConfig.Cli.Name+" version: %s\n", color.GreenString(CLI_VERSION))
	return nil
}
