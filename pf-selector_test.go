package host

import (
	"fmt"
	"testing"
)

const (
	TEST_PF_SELECTOR_VENDOR = `intel`
)

func TestPfSelectorFromVendorRegex(t *testing.T) {
	selector, err := NewPfSelector(&PfSelectorOptions{
		VendorRegexp: TEST_PF_SELECTOR_VENDOR,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = selector.SelectLocalhost()
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for _, pfs := range *selector.Selected {
		for _, _ = range pfs {
			count++
		}
	}
	if count == 0 {
		fmt.Printf("no pf selected for vendor: %s\n", TEST_PF_SELECTOR_VENDOR)
	} else {
		fmt.Printf("pfSelector number of pfs selected: %v\n", count)
	}
}
