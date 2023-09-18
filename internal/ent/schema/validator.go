package schema

import (
	"fmt"
	"net"
	"net/netip"
)

// IsValidIPAddress returns error if IP address is NOT valid
func IsValidIPAddress(ip string) error {
	if net.ParseIP(ip) != nil {
		return nil
	}

	return fmt.Errorf("Provided IP Address is invalid: %s", ip)
}

// IsValidIPPrefix returns error if IP prefix is NOT valid
func IsValidIPPrefix(prefix string) error {
	_, err := netip.ParsePrefix(prefix)
	if err != nil {
		return fmt.Errorf("Provided prefix is invalid: %s", prefix)
	}

	return nil
}
