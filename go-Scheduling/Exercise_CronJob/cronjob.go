package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	// library to interact with cron
	"github.com/go-co-op/gocron"
)

type UserData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Scheduled function
func SendRequest() {

	// Setup http client
	client := &http.Client{}

	// Setup request body
	body := map[string]interface{}{
		"name": "Jabar",
		"age":  25,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP request
	// Change the target url with your own url
	// from webhook.site
	req, err := http.NewRequest("POST", "https://webhook.site/9ef293a9-464d-4f1a-94be-0daa225435ef", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Send request
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	// Setup cron
	sched := gocron.NewScheduler(time.Local)

	// Setup scheduled request sending operation
	// similar to crontab -e
	sched.Cron("0 15 27 12 *").Do(SendRequest)

	// start cron
	sched.StartBlocking()
}
