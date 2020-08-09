// shippy/shippy-cli-consignment/main.go
package main

import (
	"fmt"
	"log"

	"context"

	pb "github.com/mindils/education-go-micro/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address1 = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address1, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	// Contact the server and print out its response.

	r, err := client.SeyHello(context.Background(), &pb.SayHelloRequest{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	fmt.Println(r.Text)

}
