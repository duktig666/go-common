// description:
// @author renshiwei
// Date: 2022/10/6 14:36

package cmd

import (
	"fmt"
	"github.com/duktig666/go-common/cmd/version"
	"github.com/duktig666/go-common/common/logger"
	"github.com/duktig666/go-common/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	cliName = config.GlobalConfig.Cli.Name
)

var rootCmd = &cobra.Command{
	Use:          cliName,
	Short:        cliName,
	SilenceUsage: true,
	Long:         cliName + `:https://github.com/duktig666/go-common`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRun: func(*cobra.Command, []string) {
		logger.InitLog(config.GlobalConfig.Log.Level, config.GlobalConfig.Log.Format)
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `Welcome to use ` + cliName + `:` + ` use ` + `-h` + ` see cli`
	usageStr1 := `You can also refer to the related content of https://github.com/duktig666/go-common`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config/config.yaml", "config file (default is config/config.yaml)")

	rootCmd.AddCommand(version.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func initConfig() {
	config.InitConfig(cfgFile)
}
