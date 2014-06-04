// +build !linux

package tuntap

import (
	"os"
)

func createInterface(f *os.File, ifPattern string, kind DevKind, flags uint16) (string, error) {
	panic("Not implemented on this platform")
}
