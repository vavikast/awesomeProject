package main

import (
	pb "awesomeProject/ES/grpc-demo/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
)
var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()

}

func main()  {
	conn, _ := grpc.Dial("192.168.14.37:"+port, grpc.WithInsecure())
	defer  conn.Close()
	client := pb.NewGreeterClient(conn)
	err :=SayHello(client)
	if err != nil {
		log.Fatalf("SayHello err: %v",err)
	}

}

func SayHello(client pb.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "eddycjy"})
	if err != nil {
		return err
	}
	log.Printf("Client.SayHello resp: %s",resp.Message)
	return  nil

}