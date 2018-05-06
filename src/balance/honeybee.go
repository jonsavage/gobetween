package balance

import (
	"fmt"

	"../core"
)

type HoneybeeBalancer struct{}

var currentBackendIndex = 0
var threshold = .8
var pollingFrequency = 5
var totalHits = 0

func (b *HoneybeeBalancer) Elect(context core.Context, backends []*core.Backend) (*core.Backend, error) {
	backendClient := core.BackendClient{}

	if shouldPoll() {
		loadOfCurrentBackend := backendClient.GetLoadForBackend(backends[currentBackendIndex])

		fmt.Println("Load of Currently Selected Backend: ", loadOfCurrentBackend)

		if loadOfCurrentBackend > threshold {
			currentBackendIndex = getIndexOfNextBackendRoundRobinStyle(backends)
			fmt.Println("Threshold exceeded, moving on to:")
			fmt.Println("\tbackend: ", currentBackendIndex)
			fmt.Println("\twith load: ", backendClient.GetLoadForBackend(backends[currentBackendIndex]))
		}
	}

	totalHits = totalHits + 1

	selectedBackend := backends[currentBackendIndex]

	return selectedBackend, nil
}

func getIndexOfNextBackendRoundRobinStyle(backends []*core.Backend) int {
	numberOfBackends := len(backends)
	indexToReturn := currentBackendIndex + 1

	if indexToReturn >= numberOfBackends {
		indexToReturn = 0
	}

	return indexToReturn
}

func shouldPoll() bool {
	return totalHits%pollingFrequency == 0
}