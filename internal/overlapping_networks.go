package internal

import (
	"net"
)

func isSubnet(sub, parent *net.IPNet) bool {
	maskedParent := parent.IP.Mask(parent.Mask)
	maskedSub := sub.IP.Mask(parent.Mask)

	pSize, _ := parent.Mask.Size()
	sSize, _ := sub.Mask.Size()

	res := maskedParent.Equal(maskedSub) && pSize <= sSize

	return res
}

func FindOverlappingNetworks(networks []*net.IPNet) map[*net.IPNet][]*net.IPNet {
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
