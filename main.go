package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

	publish(js, "test", testMsg)
	publish(js, "test.x.y.z", testXYZMsg)
	// AddConsumer(js, strName, "test", "test")
	// Update(js, strName)
	// Delete(js, strName)

	r := gin.Default()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func JetStreamContext(nc *nats.Conn) nats.JetStreamContext {
	var err error
	// nc, err = nats.Connect(nats.DefaultURL)
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

// Create a Stream
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

// Update a Stream
func Update(js nats.JetStreamContext, name string) *nats.StreamInfo {
	fmt.Printf("Update stream: %q\n", name)
	strInfo, err := js.UpdateStream(&nats.StreamConfig{
		Name:   name,
		MaxAge: 8,
	})
	if err != nil {
		log.Panicf("could not update stream: %v", err)
	}
	prettyPrint(strInfo)
	return strInfo
}

// Delete Stream
func Delete(js nats.JetStreamContext, name string) {
	fmt.Printf("Deleting stream: %q\n", name)
	if err := js.DeleteStream(name); err != nil {
		log.Printf("error deleting stream: %v", err)
	}
}

// Add Consumer
func AddConsumer(js nats.JetStreamContext, strName, consName, consFilter string) {
	info, err := js.AddConsumer(strName, &nats.ConsumerConfig{
		Durable:   consName,
		AckPolicy: nats.AckExplicitPolicy,
		// MaxAckPending: 1,      // default value is 20,000
		FilterSubject: consFilter,
	})
	if err != nil {
		log.Panicf("could not add consumer: %v", err)
	}
	prettyPrint(info)
}

func publish(js nats.JetStreamContext, subj string, f func() []byte) {
	ack, err := js.Publish(subj, f())
	if err != nil {
		log.Printf("publish error: %v", err)
	}
	fmt.Printf("%#v\n", ack)
}

func testMsg() []byte {
	return []byte(
		fmt.Sprintf("t - %s", time.Now().Format("15:04:05")),
	)
}

func testXYZMsg() []byte {
	return []byte(
		fmt.Sprintf("xyz - %s", time.Now().Format("15:04:05")),
	)
}
