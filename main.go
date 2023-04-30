package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

var deleteMode bool

func main() {
	var nc *nats.Conn
	js := JetStreamContext(nc)
	defer nc.Close()

	const strName = "tst"
	if !deleteMode {
		Create(js, strName)
		// return
	}

	// Delete(js, strName)

	r := gin.Default()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func JetStreamContext(nc *nats.Conn) nats.JetStreamContext {
	var err error
	nc, err = nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatalf("could not connect to NATS: %v", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("could not create JetStream context: %v", err)
	}

	return js
}

func Create(js nats.JetStreamContext, name string) *nats.StreamInfo {
	fmt.Printf("Creating stream: %q\n", name)
	strInfo, err := js.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{"test.>", "test"},
		MaxAge:   0,
		Storage:  nats.FileStorage,
	})
	if err != nil {
		log.Panicf("could not create stream: %v", err)
	}

	prettyPrint(strInfo)
	return strInfo
}

func prettyPrint(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatalf("could not prettyPrint: %v", err)
	}
	fmt.Println(string(b))
}

func Delete(js nats.JetStreamContext, name string) {
	fmt.Printf("Deleting stream: %q\n", name)
	if err := js.DeleteStream(name); err != nil {
		log.Printf("error deleting stream: %v", err)
	}
}
