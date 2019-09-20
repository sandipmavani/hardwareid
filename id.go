// package hardwareid provides support for reading the unique hardware id of most OSs (without admin privileges).
//
// https://github.com/sandipmavani/hardwareid
//
// https://godoc.org/github.com/sandipmavani/hardwareid/cmd/hardwareid
//
// This package is Cross-Platform (tested on Win7+, Debian 8+, Ubuntu 14.04+, OS X 10.6+, FreeBSD 11+)
// and does not use any internal hardware IDs (no MAC, BIOS, or CPU).
//
// Returned hardware IDs are generally stable for the OS installation
// and usually stay the same after updates or hardware changes.
//
// This package allows sharing of hardware IDs in a secure way by
// calculating HMAC-SHA256 over a user provided app ID, which is keyed by the hardware id.
//
// Caveat: Image-based environments have usually the same hardware-id (perfect clone).
// Linux users can generate a new id with `dbus-uuidgen` and put the id into
// `/var/lib/dbus/hardware-id` and `/etc/hardware-id`.
// Windows users can use the `sysprep` toolchain to create images, which produce valid images ready for distribution.
package hardwareid // import "github.com/sandipmavani/hardwareid"

import (
	"fmt"
	"net"
)

// ID returns the platform specific hardware id of the current host OS.
// Regard the returned id as "confidential" and consider using ProtectedID() instead.
func ID() (string, error) {
	id, err := hardwareId()
	if err != nil {
		return "", fmt.Errorf("hardwareid: %v", err)
	}
	return id, nil
}

// ProtectedID returns a hashed version of the hardware ID in a cryptographically secure way,
// using a fixed, application-specific key.
// Internally, this function calculates HMAC-SHA256 of the application ID, keyed by the hardware ID.
func ProtectedID(appID string) (string, error) {
	id, err := ID()
	if err != nil {
		return "", fmt.Errorf("hardwareid: %v", err)
	}
	return protect(appID, id), nil
}

func hardwareId() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		if i.HardwareAddr.String() != "" {
			return i.HardwareAddr.String(), nil
		}
	}

	return "", nil

}