package cmd

import (
	"fmt"
	"os"

	"github.com/david-gurley/host"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

func init() {
	rootCmd.AddCommand(newHostCmd())
}

func newHostCmd() *cobra.Command {
	hostCmd := cobra.Command{
		Use:   "host",
		Short: "display host information",
		Long:  `display host information`,
		Run: func(cmd *cobra.Command, args []string) {
			hostTable()
		},
	}
	return &hostCmd
}

func hostTable() {
	host, err := host.GetHost()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"hostname", "os", "kernel", "core", "platform", "cpus", "cpu_vendor", "cpu_model", "distribution", "release", "codename"})
	t.AppendRow([]interface{}{host.Hostname, host.OS, host.Kernel, host.Core, host.Platform, host.CPUs, host.CPUVendor, host.CPUModel, host.Distribution, host.Release, host.Codename})
	t.Render()
}
