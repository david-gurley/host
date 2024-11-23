package host

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	pfs, err := GetPfs()
	if err != nil {
		t.Fatal(err)
	}
	for _, pf := range pfs {
		//fmt.Printf("hostname: %v host_id: %v address %v\n", pf.Hostname, pf.HostID, pf.Address)
		fmt.Printf("pf: %#v\n", pf)
	}
}

func SetStuff(t *testing.T) {
	pfsMap, err := GetPfsMap()
	if err != nil {
		t.Fatal(err)
	}
	for _, pf := range pfsMap {
		//fmt.Printf("pf: %#v\n", pf)
		if pf.InterfaceName == "eno8" {
			err := SetNumVfs(pf.Address, 32)
			if err != nil {
				fmt.Printf("error setting number of vfs:  %v\n", err)
			}
			errs := pf.SetVfMacAddresses()
			if len(errs) > 0 {
				for _, err := range errs {
					fmt.Printf("error setting vf mac address: %v\n", err)
				}
			}
			//BindHostDriver(pf.Vfs)
		}
	}
}

func BindVfio(vfs Vfs) {
	for _, vf := range vfs {
		err := vf.BindVfio()
		if err != nil {
			fmt.Printf("error binding to vfio driver: %v\n", err)
		}
	}
}
func BindHostDriver(vfs Vfs) {
	for _, vf := range vfs {
		err := vf.BindHostDriver()
		if err != nil {
			fmt.Printf("error binding to host driver: %v\n", err)
		}
	}
}
