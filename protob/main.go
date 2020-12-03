package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello World!")

	elliot := &Person{
		Name: "elliot",
		Age:  24,
	}
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshalling error:", err)

	}
	fmt.Println(data)

}
