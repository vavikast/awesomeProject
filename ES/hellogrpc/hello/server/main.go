package main

import (
	pb "awesomeProject/ES/hellogrpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

const Address  = `127.0.0.1:50052`

type helloService struct {

}
var HelloService = &helloService{}

func (h *helloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message="Hello "+ req.Name+"."
	return  resp,nil
}

func main()  {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen:%v",err)
	}

	file, err := credentials.NewServerTLSFromFile("./ES/hellogrpc/keys/server.pem", "./ES/hellogrpc/keys/server.key")
	if err != nil{
		log.Fatalf("Failed to generate credentials %v",err)
	}


	//实例化grpc server
	s := grpc.NewServer(grpc.Creds(file))
	pb.RegisterHelloServer(s,HelloService)
	fmt.Println("Listen on ", Address)
	s.Serve(listen)


}