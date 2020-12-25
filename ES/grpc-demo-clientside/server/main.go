package main

import (
	pb "awesomeProject/ES/grpc-demo-clientside/proto"
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
func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.HelloReply{Message:"say.record"})

		}
		if err != nil {
			return err
		}
		log.Printf("resp: %v",resp)
	}
	return nil
}
func main()  {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server,&GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)

}