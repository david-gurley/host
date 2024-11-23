package host

import (
	"strings"
)

func GetPfByAddress(address string) (*Pf, error) {
	return &Pf{}, nil
}

func GetPfs() ([]*Pf, error) {
	pfs := make([]*Pf, 0)
	return pfs, nil
}

// should we look at getting information fron netlink vs. traversing
// directory structures?
func GetPf(address string) (Pf, error) {
	var pf Pf
	return pf, nil
}

func IsPf(address string) bool {
	return false
}

func (pf *Pf) Stats() (map[string]uint64, error) {
	stats := make(map[string]uint64)
	return stats, nil
}

func (pf *Pf) Features() (map[string]bool, error) {
	features := make(map[string]bool)
	return features, nil
}
func (pf *Pf) LinkState() (uint32, error) {
	return 0, nil
}

func (pf *Pf) GetVfsConfig() error {
	return nil
}

// apply the pf-policy concrete config
// we don't want to return if there is
// a single error - return slice of errors
func ApplyPfConfigs(pfConfigs *[]PfConfig) []error {
	errs := make([]error, 0)
	return errs
}

func SetNumVfs(address string, num int) error {
	return nil
}

func (pf *Pf) SetNumVfs(num int) error {
	return nil
}

func (pf *Pf) PopulateLinkAttrs() error {
	return nil
}

func (pf *Pf) IPAddressesJoin() string {
	return strings.Join(pf.IPAddresses, ",")
}

// handle (socket) for the network requests on a specific network namespace.
func (pf *Pf) SetVfMacAddresses() []error {
	errs := make([]error, 0)
	return errs
}

func (pf *Pf) SetVfMacAddressesNetlink() error {
	return nil
}

func (pf *Pf) LinkSetUp() error {
	return nil
}

func PfFromVfAddress(address string) Pf {
	return Pf{}
}

func (pf *Pf) PopulateVfIndexes() error {
	return nil
}

func PopulateVfIndexes(pf *Pf) error {
	return nil
}
