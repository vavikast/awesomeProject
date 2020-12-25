package main

import (
	pb "awesomeProject/ES/grpc-demo-serverside/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"io"
	"log"
)
var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()

}

func main()  {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer  conn.Close()
	client := pb.NewGreeterClient(conn)
	err :=SayHello(client,&pb.HelloRequest{
		Name:                 "eddyjj",
	})
	if err != nil {
		log.Fatalf("SayHello err: %v",err)
	}

}

func SayHello(client pb.GreeterClient,r *pb.HelloRequest) error {
	stream, err := client.SayList(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			return  err
		}
		log.Printf("resp: %v",resp)
	}
	return nil
}