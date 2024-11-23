package cmd

import (
	"fmt"
	"os"

	"github.com/david-gurley/host"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

func init() {
	rootCmd.AddCommand(newPfCmd())
}

func newPfCmd() *cobra.Command {
	pfCmd := cobra.Command{
		Use:   "pf",
		Short: "display physical functions",
		Long:  `display physical functions`,
		Run: func(cmd *cobra.Command, args []string) {
			pfTable()
		},
	}
	//registerCmd.Flags().StringVarP(&accountID, "account-id", "a", "", "account id for agent registration")
	return &pfCmd
}

func pfTable() {
	pfs, err := host.GetPfs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"address", "vendor", "device", "interface_name", "mac_address", "driver", "driver_version",
		"fw_version", "total_vfs", "num_vfs", "has_ip_address", "bond_member", "has_subinterfaces",
	})
	for _, pf := range pfs {
		t.AppendRow([]interface{}{
			pf.Address, pf.Vendor, pf.Device, pf.InterfaceName, pf.MacAddress, pf.Driver, pf.DriverVersion,
			pf.FwVersion, pf.TotalVfs, pf.NumVfs, pf.HasIPAddress, pf.BondMember, pf.HasSubinterfaces,
		})
	}
	t.Render()
}
