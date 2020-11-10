package extensionsapiclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Register creates registration of lambda extension
func Register(agentName string, registrationBody interface{}) (interface{}, error) {
	runtimeAPIAddress, exists := os.LookupEnv("AWS_LAMBDA_RUNTIME_API")
	if !exists {
		return nil, errors.New("AWS_LAMBDA_RUNTIME_API is not set")
	}

	registrationBodyJSON, err := json.Marshal(registrationBody)
	if err != nil {
		return nil, err
	}

	log.Println("Registering to Extensions API")

	client := &http.Client{}
	request, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("http://%s/2020-01-01/extension/register", runtimeAPIAddress),
		bytes.NewBuffer(registrationBodyJSON),
	)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Lambda-Extension-Name", agentName)
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
		log.Fatalln("Could not register to Extensions API:", response.StatusCode, responseBody)
	}

	return response.Header.Get("Lambda-Extension-Identifier"), nil
}
