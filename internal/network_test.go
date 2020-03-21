package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPNetworkContains(t *testing.T) {
	network := NewIPV4Network("")
	subNetwork := NewIPV4Network("")
	subSubNetwork := NewIPV4Network("")

	assert.True(t, network.Contains(subNetwork))
	assert.True(t, network.Contains(subSubNetwork))

	assert.False(t, subNetwork.Contains(network))
	assert.True(t, subNetwork.Contains(subSubNetwork))

	assert.False(t, subSubNetwork.Contains(network))
	assert.False(t, subSubNetwork.Contains(subNetwork))
}
