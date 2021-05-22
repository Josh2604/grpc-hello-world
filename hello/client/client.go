package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Josh2604/grpc-hello-world/hello"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client running!")
	// Por default los servicios de RPC son seguros por lo tanto requieren de un certificado
	// para omitirlo agregamos grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	// Implementación de los metodos definidos mediante el proto
	c := hello.NewHelloServiceClient(conn)
	doHello(c)
}

func doHello(c hello.HelloServiceClient) {
	fmt.Println("Starting rpc")
	req := &hello.HelloRequest{
		User: &hello.User{
			Name: "test",
			Age:  "25",
		},
	}
	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling rpc method: %v", err)
	}

	log.Printf("Response from GreatService: %v", res.Message)
}
