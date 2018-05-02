package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type load struct {
	Number float64 `json:"load"`
}

type BackendClient struct{}

func (backendClient *BackendClient) GetLoadForBackend(backend *Backend) float64 {

	restClient := http.Client{}
	url := "http://" + backend.Target.Address() + "/load"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error!", err)
	}

	res, getErr := restClient.Do(request)
	if getErr != nil {
		fmt.Println("error!", getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("error!", readErr)
	}

	load := load{}
	jsonErr := json.Unmarshal(body, &load)
	if jsonErr != nil {
		fmt.Println("jsonErr", jsonErr)
	}

	return load.Number
}
