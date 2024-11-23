package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "slickhost",
		Short: "slick host utilities",
		Long:  `slick host utilities`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
