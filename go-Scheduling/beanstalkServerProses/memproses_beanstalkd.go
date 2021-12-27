package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	bs "github.com/beanstalkd/go-beanstalk"
)

// struct to receive data from queue system
type UserData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("this is processing server")
	// Setup connection to Beanstalkd
	conn, err := bs.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		log.Fatal(err)
	}

	// Pull data from Beanstalkd
	id, body, err := conn.Reserve(5 * time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	// Assign data to variable for processing purpose
	var data UserData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Processing data
	log.Println("Processing data!")
	log.Println("Data: ", data)

	// Removing data from queue system
	// so it'll be replaced by the second oldest
	// data there
	err = conn.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
}
