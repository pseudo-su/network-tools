package internal

func FindOverlappingNetworks(networks []*IPV4Network) map[*IPV4Network][]*IPV4Network {
	result := map[*IPV4Network][]*IPV4Network{}

	for i, network1 := range networks {
		for j := i + 1; j < len(networks); j++ {
			network2 := networks[j]
			switch {
			case network1.Contains(network2):
				if _, ok := result[network1]; !ok {
					result[network1] = []*IPV4Network{}
				}
				result[network1] = append(result[network1], network2)
			case network2.Contains(network1):
				if _, ok := result[network1]; !ok {
					result[network1] = []*IPV4Network{}
				}
				result[network2] = append(result[network2], network1)
			}
		}
	}

	return result
}

type NetworkOverlapResult int

const (
	NoOverlap NetworkOverlapResult = iota
	FirstOverlapsSecond
	SecondOverlapsFirst
)

func isOverlapping(net1, net2 *IPV4Network) NetworkOverlapResult {
	return NoOverlap
}
