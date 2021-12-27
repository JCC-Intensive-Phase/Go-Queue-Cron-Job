package main

import (
	"encoding/json"
	"log"
	"time"

	// library to interact with Beanstalkd
	bs "github.com/beanstalkd/go-beanstalk"
	// library for http server
	"github.com/gin-gonic/gin"
)

type UserData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	log.Println("This is receiver server")
	log.Println("Setting up connection with Beanstalkd")
	// Setup connection to Beanstalkd
	conn, err := bs.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		log.Fatal(err)
	}

	// Setup http server
	r := gin.Default()
	// Setup endpoint to receive user data
	r.POST("/user", func(c *gin.Context) {

		var data UserData
		// Read HTTP request body and bind it to 'data'
		err := c.BindJSON(&data)
		if err != nil {
			log.Fatal(err)
		}

		// Marshal data to []byte to be put to Beanstalkd
		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Put(dataBytes, 1, 0, 120*time.Minute)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(202, gin.H{
			"message": "processing user data",
		})
	})
	r.Run(":8080")

}
