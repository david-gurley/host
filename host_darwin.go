package host

import (
	"runtime"
)

const (
	hostIdPath = "/etc/machine-id"
)

func GetHost() (*Host, error) {
	var host Host
	host.OS = runtime.GOOS
	host.CPUs = runtime.NumCPU()
	hostname, err := GetHostname()
	if err != nil {
		return &host, err
	}
	host.Hostname = hostname
	return &host, nil
}

func (host *Host) updateUname() error {
	return nil
}

func (host *Host) updateLsbRelease() error {
	return nil
}

func GetHostID() (string, error) {
	return ReadFileString(hostIdPath)
}

func GetKernelCmdline() (string, error) {
	return "", nil
}
