package schema

import (
	"context"
	"fmt"
	"net"

	"go.infratographer.com/x/gidx"
	// "go.infratographer.com/x/gidx"
)

// IsValidIPAddress returns true if IP address is a valid or not
//
//	for _, edge := range IPAddresses.Edges {
//		ipAdr := edge.Node.IP
//	}
func IsValidIPAddress(ip string) error {
	if net.ParseIP(ip) != nil {
		return nil
	}

	return fmt.Errorf("Provided IP Address is invalid: %s", ip)
}

// IsValidIPPrefix validates if the prefix of the IP is valid or not
// IPBlock.Prefix
func IsValidIPPrefix(prefix string) (bool, error) {
	// TODO: once new version of gidx is loaded, uncomment below
	// if gidx.ValidPrefix(prefix) {
	// 	return true, nil
	// }

	return false, fmt.Errorf("Provided prefix is invalid: %s", prefix)
}

// - Valid IP outside of block is rejected
// - Valid IP inside of block is accepted

// IsPartOfBlock validates if IP is part of associated block
func IsPartOfBlock(ip string, block string) (bool, error) {

	ct, dr, err := OpenTxFromContext()
	ipBlock, err := dr.GetIPBlock(ctx, block)

	ipBlock.IPAddresses

	for _, address := range ipBlock.IPAddresses {
		if address.IP == ip {
			return true, nil
		}
	}
	return false, fmt.Errorf("Provided IP is not part of associated block: %s", ip)
}

// getIPBlock returns an IP Block by id
func (c *Client) getIPBlock(ctx context.Context, id string) (*GetIPBlockResult, error) {
	_, err := gidx.Parse(id)
	if err != nil {
		return nil, err
	}

	vars := map[string]interface{}{
		"id": id,
	}

	var ipb GetIPBlockResult
	if err := c.gqlCli.Query(ctx, &ipb, vars); err != nil {
		return nil, err
	}

	return &ipb, nil
}
