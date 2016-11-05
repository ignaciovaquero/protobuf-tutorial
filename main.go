package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/igvaquero18/protobuf-tutorial/tutorial"
)

func main() {
	person := &pb.Person{
		Name:  "John Appleseed",
		Id:    1,
		Email: "johnappleseed@email.com",
		Phones: []*pb.Person_PhoneNumber{
			&pb.Person_PhoneNumber{
				Number: "612345678",
				Type:   pb.Person_MOBILE,
			},
			&pb.Person_PhoneNumber{
				Number: "912345678",
				Type:   pb.Person_HOME,
			},
		},
	}

	b, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Error when encoding: %s\n", err.Error())
	}

	fmt.Printf("Encoded person: %v\n", b)
	fmt.Printf("Person: %s\n", string(b))
}
