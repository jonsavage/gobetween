package balance

import (
	"fmt"
	"math/rand"
	"../core"
)

type HoneybeeBalancer struct{}

var currentBackendIndex = 0
var threshold = 50
var pollingFrequency = 5
var totalHits = 0

func (b *HoneybeeBalancer) Elect(context core.Context, backends []*core.Backend) (*core.Backend, error) {
	backendClient := core.BackendClient{}

	if shouldPoll() {
		loadOfCurrentBackend := backendClient.GetLoadForBackend(backends[currentBackendIndex])

		fmt.Println("Load of Currently Selected Backend: ", loadOfCurrentBackend)

		if loadOfCurrentBackend > threshold {
			currentBackendIndex = getIndexOfNextBackendHoneyBeeStyle(backends)
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

func getIndexOfNextBackendHoneyBeeStyle(backends []*core.Backend) int {
	backendClient := core.BackendClient{}
	//init smallest load with first backends load
	var smallestLoad = backendClient.GetLoadForBackend(backends[0])
	var smallestLoadIndex = 0
	for index := range backends {
		loadForBackend := backendClient.GetLoadForBackend(backends[index])
		fmt.Println("load For: ", index)
		fmt.Println(loadForBackend)

		if loadForBackend < smallestLoad {
			smallestLoad = loadForBackend
			smallestLoadIndex = index
		}
	}
	return smallestLoadIndex
}

func getIndexForNextBackendRandomStyle(backends []*core.Backend) int {
	randomIndex := rand.Intn(len(backends))

	return randomIndex
}

// myrand := random(1, 6)

func shouldPoll() bool {
	return totalHits%pollingFrequency == 0
}
