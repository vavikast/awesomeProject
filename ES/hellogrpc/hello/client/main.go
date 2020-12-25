package main

import (
	pb "awesomeProject/ES/hellogrpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const  Address  = "127.0.0.1:50052"

func main()  {
	creds, err := credentials.NewClientTLSFromFile("./ES/hellogrpc/keys/server.pem", "TXXY")
	if err != nil{
		log.Fatalf("Failed to create tls credentials %v",err)
	}



	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)
	reqBody.Name= "grpc"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r.Message)

}