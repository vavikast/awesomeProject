package main

import (
	pb "awesomeProject/ES/grpc-demo/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()

}

type GreeterServer struct {
}

func (s *GreeterServer)SayHello(ctx context.Context,r *pb.HelloRequest)(*pb.HelloReply,error)  {
	return &pb.HelloReply{Message: "Hello.world"},nil
}
func main()  {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server,&GreeterServer{})
	lis, _ := net.Listen("tcp", "192.168.14.37:"+port)
	server.Serve(lis)

}