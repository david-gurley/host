package host

import (
	"fmt"
	"testing"
)

func TestPfPolicy(t *testing.T) {
	pfPolicy, err := NewPfPolicy(
		&PfSelectorOptions{
			VendorRegexp: TEST_PF_SELECTOR_VENDOR,
		},
		&PfPolicyOptions{
			OnReboot: true,
			MaxVfs:   true,
		})
	if err != nil {
		t.Fatal(err)
	}
	err = pfPolicy.PfSelector.SelectLocalhost()
	if err != nil {
		t.Fatal(err)
	}
	err = pfPolicy.Plan()
	if err != nil {
		t.Fatal(err)
	}
	for hostID, pfMap := range pfPolicy.PfPlan {
		for pfID, kindMap := range pfMap {
			fmt.Printf("host_id: %s pci_address: %s current num_vfs: %v planned num_vfs: %v\n",
				hostID, pfID, kindMap["current"].NumVfs, kindMap["planned"].NumVfs)
		}
	}
}
