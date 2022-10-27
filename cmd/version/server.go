package version

import (
	"fmt"
	"github.com/duktig666/go-common/common/global"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "couples-cli version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Printf("couples-subtotal version: %s\n", color.GreenString(global.VERSION))
	return nil
}
