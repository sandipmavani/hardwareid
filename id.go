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
	"crypto/md5"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
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

func physicalId() (string, error) {
	id, err := GetPhysicalId()
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
func Md5(data []byte) string {
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%X", has)
	return md5Str
}
func GetPhysicalId() (string, error) {
	biosSn, err := GetBIOSSerialNumber()
	if err != nil {
		return "", err
	}
	diskSn, err := GetDiskDriverSerialNumber()
	if err != nil {
		return "", err
	}
	cpuId, err := GetCPUPorcessorID()
	if err != nil {
		return "", err
	}
	return strings.ToUpper(Md5([]byte(biosSn + diskSn + cpuId))), nil
}

//获取硬盘SerialNumber
func GetDiskDriverSerialNumber() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return diskDriverSerialNumberOnWindows()
	case "linux":
		return "", errors.New("unsupported os")
	case "darwin":
		return "", errors.New("unsupported os")
	default:
		return "", errors.New("unknown os")
	}
}

func diskDriverSerialNumberOnWindows() (string, error) {
	cmd := exec.Command("CMD", "/C", "WMIC DISKDRIVE GET SERIALNUMBER")
	serialNo, err := cmd.Output()
	if err != nil {
		return "", err
	}
	l := strings.Split(string(serialNo), "\n")
	if len(l) >= 2 {
		return l[1], nil
	} else {
		return "", errors.New("return split length less 2")
	}
}

//获取硬盘SerialNumber
func GetBIOSSerialNumber() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return biosSerialNumberOnWindows()
	case "linux":
		return "", errors.New("unsupported os")
	case "darwin":
		return "", errors.New("unsupported os")
	default:
		return "", errors.New("unknown os")
	}
}

func biosSerialNumberOnWindows() (string, error) {
	cmd := exec.Command("CMD", "/C", "WMIC BIOS GET SERIALNUMBER")
	serialNo, err := cmd.Output()
	if err != nil {
		return "", err
	}
	l := strings.Split(string(serialNo), "\n")
	if len(l) >= 2 {
		newline := strings.Split(l[1], " ")
		return newline[0], nil
	} else {
		return "", errors.New("return split length less 2")
	}
}

//获取CPU ProcessorID
func GetCPUProcessorID() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return cpuPorcessorIDOnWindows()
	case "linux":
		return "", errors.New("unsupported os")
	case "darwin":
		return "", errors.New("unsupported os")
	default:
		return "", errors.New("unknown os")
	}
}

func cpuProcessorIDOnWindows() (string, error) {
	cmd := exec.Command("CMD", "/C", "WMIC CPU GET ProcessorID")
	serialNo, err := cmd.Output()
	if err != nil {
		return "", err
	}
	l := strings.Split(string(serialNo), "\n")
	if len(l) >= 2 {
		newline := strings.Split(l[1], " ")
		return newline[0], nil
	} else {
		return "", errors.New("return split length less 2")
	}
}
