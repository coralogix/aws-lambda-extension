package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/coralogix/aws-lambda-extension/pkg/coralogixapiclient"
	"github.com/coralogix/aws-lambda-extension/pkg/extensionsapiclient"
	"github.com/coralogix/aws-lambda-extension/pkg/httplistener"
	"github.com/coralogix/aws-lambda-extension/pkg/logsapiclient"
)

func main() {
	agentName := path.Base(os.Args[0])
	listenerState := make(chan bool)
	queue := make(chan interface{})

	log.Println("Initializing Lambda Extension", agentName)
	agentID, err := extensionsapiclient.Register(agentName, map[string]interface{}{
		"events": []string{"INVOKE", "SHUTDOWN"},
	})
	if err != nil {
		log.Fatalln("Failed to register Lambda Extension", agentName)
	}

	go httplistener.Serve(queue, listenerState)
	select {
	case <-listenerState:
		logsapiclient.Subscribe(agentID.(string), map[string]interface{}{
			"destination": map[string]interface{}{
				"protocol": "HTTP",
				"URI":      fmt.Sprintf("http://sandbox:%d", httplistener.ReceiverPort),
			},
			"types": []string{"platform", "function"},
			"buffering": map[string]uint{
				"timeoutMs": 1000,
				"maxBytes":  1048576,
				"maxItems":  10000,
			},
		})

		for {
			extensionsapiclient.Next(agentID.(string))
			coralogixapiclient.Send((<-queue).([]interface{}))
		}
	case <-time.After(9 * time.Second):
		log.Fatalln("HTTP Server has timedout before starting")
	}
}
