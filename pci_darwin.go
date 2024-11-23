package host

func GetPciDeviceClass(address string) (string, error) {
	return "", nil
}
func GetPciDeviceVendor(address string) (string, error) {
	return "", nil
}
func GetPciDeviceDevice(address string) (string, error) {
	return "", nil
}

func InterfaceNameFromAddress(address string) string {
	return ""
}

// is the pci device controlled by vfio?
func IsVfio(address string) bool {
	return false
}

// get the iommu group for a given pci bus address
func IommuGroup(address string) string {
	return ""
}

// get the driver for a give pci bus address
func Driver(address string) string {
	return ""
}

// is the given iommu group allocated to a process?
func IsAllocated(allocations []string, iommuGroup string) bool {
	return false
}

// Find all current vfio allocations
// part of this from mitchellh/go-ps/process_unix.go
func VfioAllocations() ([]string, error) {
	allocations := make([]string, 0)
	return allocations, nil
}
