package extensionsapiclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Next pulls lambda extension next event
func Next(agentID string) (interface{}, error) {
	runtimeAPIAddress, exists := os.LookupEnv("AWS_LAMBDA_RUNTIME_API")
	if !exists {
		return nil, errors.New("AWS_LAMBDA_RUNTIME_API is not set")
	}

	client := &http.Client{}
	request, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("http://%s/2020-01-01/extension/event/next", runtimeAPIAddress),
		nil,
	)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Lambda-Extension-Identifier", agentID)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	responseBody := string(responseBytes)

	if response.StatusCode != 200 {
		log.Fatalln("Request to Extensions API failed:", response.StatusCode, responseBody)
	}

	//log.Println("Received response from Extensions API:", responseBody)

	return responseBody, nil
}
