package main

import (
	pb "awesomeProject/ES/grpc-demo-bid/proto"
	"flag"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()

}

type GreeterServer struct {
}

func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.HelloReply{Message:"say.route"})
		if err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}
		n++
		log.Printf("resp: %v,number: %s",resp,n)
	}
}
func main()  {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server,&GreeterServer{})
	lis, _ := net.Listen("tcp", "127.0.0.1:"+port)
	server.Serve(lis)

}