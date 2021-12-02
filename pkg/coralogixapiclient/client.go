package coralogixapiclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Send builds logs batches and send to Coralogix
func Send(functionName string, records []interface{}) {
	coralogixURL := GetEnv("CORALOGIX_LOG_URL", "https://api.coralogix.com/api/v1/logs")
	privateKey := GetEnv("CORALOGIX_PRIVATE_KEY", "")
	applicationName := GetEnv("CORALOGIX_APP_NAME", functionName)
	subsystemName := GetEnv("CORALOGIX_SUB_SYSTEM", "logs")
	logEntries := []map[string]interface{}{}

	if privateKey == "" {
		log.Fatalln("CORALOGIX_PRIVATE_KEY is not set")
	}

	if len(records) > 0 {
		for _, record := range records {
			var text string
			record := record.(map[string]interface{})
			timestamp, _ := time.Parse("2006-01-02T15:04:05.000Z", record["time"].(string))

			switch v := record["record"].(type) {
			case string:
				text = string(v)
			default:
				jsonText, _ := json.Marshal(v)
				text = string(jsonText)
			}

			logEntries = append(logEntries, map[string]interface{}{
				"timestamp": timestamp.UnixNano() / 1000000,
				"severity":  GetSeverityLevel(text),
				"text":      text,
				"category":  record["type"],
			})
		}

		logsBulk, _ := json.Marshal(map[string]interface{}{
			"privateKey":      privateKey,
			"applicationName": applicationName,
			"subsystemName":   subsystemName,
			"logEntries":      logEntries,
		})

		client := &http.Client{}
		request, _ := http.NewRequest("POST", coralogixURL, bytes.NewBuffer(logsBulk))
		request.Close = true
		request.Header.Set("Content-Type", "application/json")
		response, err := client.Do(request)
		if err != nil {
			log.Println("Cannot send logs to Coralogix:", err)
		} else {
			defer response.Body.Close()
			if response.StatusCode != 200 {
				log.Println("Coralogix API returned unsuccess code:", response.StatusCode)
			}
		}
	}
}
