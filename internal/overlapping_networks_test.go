package internal

import (
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CIDRFixture(t *testing.T, in string) *net.IPNet {
	_, nw, err := net.ParseCIDR(in)
	require.NoError(t, err)
	return nw
}

func TestFindOverlappingNetworks(t *testing.T) {
	fixtureString := `
192.168.0.0/16
192.168.1.0/24
192.168.1.64/29
192.168.5.54/32
192.168.52.0/20
10.0.50.0/24
10.0.0.0/28
10.0.0.0/8
`
	networks, err := ReadNetworks(strings.NewReader(fixtureString))
	require.NoError(t, err)
	overlapping := FindOverlappingNetworks(networks)

	assert.NotEmpty(t, overlapping)
}

func TestIsSubnet(t *testing.T) {
	_, network, err := net.ParseCIDR("192.168.0.0/16")
	require.NoError(t, err)
	_, subNetwork, err := net.ParseCIDR("192.168.1.0/24")
	require.NoError(t, err)
	_, subSubNetwork, err := net.ParseCIDR("192.168.1.64/29")
	require.NoError(t, err)

	assert.True(t, isSubnet(subNetwork, network))
	assert.True(t, isSubnet(subSubNetwork, network))

	assert.False(t, isSubnet(network, subNetwork))
	assert.True(t, isSubnet(subSubNetwork, subNetwork))

	assert.False(t, isSubnet(network, subSubNetwork))
	assert.False(t, isSubnet(subNetwork, subSubNetwork))
}
