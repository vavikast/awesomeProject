package main

import (
	pb "awesomeProject/ES/grpc-demo-serverside/proto"
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
func (s *GreeterServer) SayList(req *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n :=0;n<=6;n++{
		err:= stream.Send(&pb.HelloReply{
			Message:"hello.list"})
		if err != nil {
			return err
		}
	}
	return  nil

}
func main()  {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server,&GreeterServer{})

	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)

}