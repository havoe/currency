package cmd

import (
	"github.com/spf13/cobra"
	"os"
)


var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "currency",
		Short: "currency Api server",
		Long: `start currency Api server`,
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	}
)

//Execute : run commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	rootCmd.AddCommand(StartCmd)
}

