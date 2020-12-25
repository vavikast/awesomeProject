package main

import (
	pb "awesomeProject/ES/grpc-demo-clientside/proto"
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
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer  conn.Close()
	client := pb.NewGreeterClient(conn)
	err :=SayHello(client,&pb.HelloRequest{
		Name:                 "You are beautiful",
	})
	if err != nil {
		log.Fatalf("SayHello err: %v",err)
	}

}

func SayHello(client pb.GreeterClient,r *pb.HelloRequest) error {
	stream, err := client.SayRecord(context.Background())
	if err != nil {
		return err
	}
	 for n := 0; n<6;n++ {
		 err := stream.Send(r)
		 if err != nil {
			 return err
		 }
	 }
	return nil
}