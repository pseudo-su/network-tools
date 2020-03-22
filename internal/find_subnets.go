package internal

import (
	"net"
)

func isSubnet(possibleSubnet, possibleSupernet *net.IPNet) bool {
	maskedSupernet := possibleSupernet.IP.Mask(possibleSupernet.Mask)
	maskedSub := possibleSubnet.IP.Mask(possibleSupernet.Mask)

	superSize, _ := possibleSupernet.Mask.Size()
	subSize, _ := possibleSubnet.Mask.Size()

	res := maskedSupernet.Equal(maskedSub) && superSize <= subSize

	return res
}

func FindSubnets(networks []*net.IPNet) map[*net.IPNet][]*net.IPNet {
	result := map[*net.IPNet][]*net.IPNet{}

	for i, network1 := range networks {
		for j := i + 1; j < len(networks); j++ {
			network2 := networks[j]
			switch {
			case isSubnet(network2, network1):
				if _, ok := result[network1]; !ok {
					result[network1] = []*net.IPNet{}
				}
				result[network1] = append(result[network1], network2)
			case isSubnet(network1, network2):
				if _, ok := result[network1]; !ok {
					result[network1] = []*net.IPNet{}
				}
				result[network2] = append(result[network2], network1)
			}
		}
	}

	return result
}
