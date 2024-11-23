package cmd

import (
	"fmt"
	"os"

	"github.com/david-gurley/host"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

func init() {
	rootCmd.AddCommand(newVfCmd())
}

func newVfCmd() *cobra.Command {
	vfCmd := cobra.Command{
		Use:   "vf",
		Short: "display virtual functions",
		Long:  `display virtual functions`,
		Run: func(cmd *cobra.Command, args []string) {
			vfTable()
		},
	}
	//registerCmd.Flags().StringVarP(&accountID, "account-id", "a", "", "account id for agent registration")
	return &vfCmd
}

func vfTable() {
	vfs, err := host.GetVfs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(vfs) == 0 {
		fmt.Println("no vfs found on system")
		os.Exit(0)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"allocated", "address", "vendor", "interface_name", "mac_address", "ip_addresses", "device", "driver", "pf_address"})
	for _, vf := range vfs {
		t.AppendRow([]interface{}{vf.Allocated, vf.Address, vf.Vendor, vf.InterfaceName, vf.MacAddress, vf.IPAddressesJoin(), vf.Device, vf.Driver, vf.PfAddress})
	}
	t.Render()
}
