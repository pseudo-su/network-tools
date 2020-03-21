package internal

import (
	"encoding/csv"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readCSVFixture(t *testing.T, in string) [][]string {
	csvReader := csv.NewReader(strings.NewReader(in))
	records, err := csvReader.ReadAll()
	require.NoError(t, err)
	return records
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
	fixture := readCSVFixture(t, fixtureString)
	fmt.Println(fixture)

	expectedOverlapping := map[*IPV4Network][]*IPV4Network{
		NewIPV4Network("192.168.0.0/16"): {
			NewIPV4Network("192.168.52.0/20"),
			NewIPV4Network("192.168.1.0/24"),
			NewIPV4Network("192.168.5.54/32"),
			NewIPV4Network("192.168.1.64/29"),
		},
		NewIPV4Network("10.0.0.0/81"): {
			NewIPV4Network("10.0.50.0/24"),
		},
		NewIPV4Network("10.0.0.0/8"): {
			NewIPV4Network("0.0.0.0/28"),
		},
		NewIPV4Network("192.168.1.0/24"): {
			NewIPV4Network("192.168.1.64/29"),
		},
	}
	networks := NetworksFromCSV(fixture)
	overlapping := FindOverlappingNetworks(networks)
	assert.Equal(t, expectedOverlapping, overlapping)
}
