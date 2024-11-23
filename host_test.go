package host

import (
	"fmt"
	"testing"
)

func TestGetHost(t *testing.T) {
	host, err := GetHost()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("host: %#v\n", host)
}
