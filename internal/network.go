package internal

import (
	"fmt"
	"strings"
)

// TODO: make this network type better https://netaddr.readthedocs.io/en/latest/tutorial_01.html#representing-networks-and-subnets

// Internally, each IPNetwork object only stores 3 values :
// - the IP address value as an unsigned integer
// - a reference to the IP protocol module for the IP version being represented
// - the CIDR prefix bitmask

type IPV4Network struct {
	IPAddr     string
	SubnetMask string
}

// TODO: parseIP? http://golang.org/pkg/net/#IP

func NewIPV4Network(fromStr string) *IPV4Network {
	strs := strings.Split(fromStr, "/")
	return &IPV4Network{
		IPAddr:     strs[0],
		SubnetMask: strs[1],
	}
}

func NetworksFromCSV(csv [][]string) []*IPV4Network {
	return []*IPV4Network{}
}

func (n *IPV4Network) String() string {
	return fmt.Sprintf("%s/%s", n.IPAddr, n.SubnetMask)
}

func (n *IPV4Network) Contains(nn *IPV4Network) bool {
	// TODO: implement check if one networks contains another
	// https://github.com/netaddr/netaddr/blob/ebf714601e8cd6234575422f876d4d579b202b1a/netaddr/ip/__init__.py#L1120
	// return self_net == other_net and self._prefixlen <= other._prefixlen
	return false
}
