package host

func UnbindDriver(vfAddress string) error {
	return nil
}

func (vf *Vf) UnbindDriver() error {
	return nil

}

func (vf *Vf) BindVfio() error {
	return nil
}

func (vf *Vf) BindHostDriver() error {
	return nil
}

func (vfs *Vfs) GetAllocations() error {
	return nil
}

func GetVfs() (Vfs, error) {
	return Vfs{}, nil
}

func GetVf(address string) (Vf, error) {
	return Vf{}, nil
}

func VfPf(address string) Pf {
	return Pf{}
}

func (vf *Vf) GetMacAddress() error {
	return nil
}

func (vf *Vf) GetIPAddresses() error {
	return nil
}

func (vf *Vf) IPAddressesJoin() string {
	return ""
}
