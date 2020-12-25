package main

import (
	pb "awesomeProject/ES/grpc-demo-bid/proto"
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
		Name:                 "eddycty",
	})
	if err != nil {
		log.Fatalf("SayHello err: %v",err)
	}

}

func SayHello(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayRoute(context.Background())
	if err != nil {
		return err
	}
	for n:=0;n<=6;n++{
		err := stream.Send(r)
		if err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err==io.EOF{
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp err: %v",resp)
	}
	err = stream.CloseSend()
	if err != nil {
		return err
	}
	return nil
}