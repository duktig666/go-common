// description:
// @author renshiwei
// Date: 2022/10/6 14:36

package cmd

import (
	"fmt"
	"github.com/duktig666/go-common/cmd/version"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "go-common",
	Short:        "go-common",
	SilenceUsage: true,
	Long:         `go-common:https://github.com/duktig666/go-common`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `Welcome to use ` + `go-common:` + ` use ` + `-h` + ` see cli`
	usageStr1 := `You can also refer to the related content of https://github.com/duktig666/go-common 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(version.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
