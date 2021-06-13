package API

import (
	"fmt"
	"github.com/achenet/TAT/request_reader"
	"net/http"
)

type API struct {
	protocol string
	hostname string
}

func (api *API) SendRequest(request *request_reader.Request) (err error) {
	targetEndpoint := api.protocol + api.hostname + request.Endpoint

	client := &http.Client{}

	httpRequest, err := http.NewRequest(request.Method, targetEndpoint, request.Data)
	if err != nil {
		return err
	}
	for _, header := range api.Headers {
		httpRequest.Header.Add(header.Key, header.Value)
	}

	resp, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
