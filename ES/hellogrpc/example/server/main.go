package main

import (
	pb "awesomeProject/ES/hellogrpc/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

const Address  = "127.0.0.1:50052"


type helloService struct {

}
var HelloService = &helloService{}

func (h *helloService)SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	
	resp := new(pb.HelloReply)
	resp.Message= "Hello "+ req.Name+"."
	return resp,nil

}

func auth(ctx  context.Context) error  {
	md, ok := metadata.FromIncomingContext(ctx)
	
	if !ok {
		return errors.New("无Token认证信息")
	}
	var (
		appid string
		appkey string
	)
	if val,ok := md["appid"];ok{
		appid = val[0]
	}
	if val,ok := md["appkey"];ok{
		appkey = val[0]
	}
	if appid != "101010" || appkey !="i am key"{
		return errors.New("Token 认证信息无效")
	}
	return nil
}
func main()  {
	var opts []grpc.ServerOption
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}
	creds, err := credentials.NewServerTLSFromFile("./ES/hellogrpc/keys/server.pem", "./ES/hellogrpc/keys/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v",err)
	}

	opts = append(opts,grpc.Creds(creds))
	
	//注册interceptor
	var interceptor  grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = auth(ctx)
		if err != nil {
			return
		}
		return handler(ctx,req)
	}

	opts = append(opts,grpc.UnaryInterceptor(interceptor))
	s := grpc.NewServer(opts...)
	pb.RegisterHelloServer(s,HelloService)

	fmt.Println("Listen on "+Address+" with TLS Token")
	s.Serve(listen)

}