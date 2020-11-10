package httplistener

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	// ReceiverHost describes logs receiver hostname
	ReceiverHost = "0.0.0.0"
	// ReceiverPort describes logs receiver port
	ReceiverPort = 4342
)

// Serve start HTTP server to accept incomming events
func Serve(queue chan<- interface{}, state chan bool) {
	address := fmt.Sprintf("%s:%d", ReceiverHost, ReceiverPort)

	log.Println("Initializing HTTP Server on", address)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			var contentLength int
			var requestBody interface{}

			contentLength, err := strconv.Atoi(request.Header.Get("Content-Length"))
			if err != nil {
				contentLength = 0
			}

			if contentLength > 0 {
				defer request.Body.Close()
				requestBytes, err := ioutil.ReadAll(request.Body)
				if err != nil {
					panic(err)
				}

				err = json.Unmarshal(requestBytes, &requestBody)
				if err != nil {
					panic(err)
				}

				queue <- requestBody
			}

			writer.WriteHeader(http.StatusOK)
			return
		}
	})

	log.Println("Serving HTTP Server on", address)
	state <- true
	http.ListenAndServe(address, nil)
}
