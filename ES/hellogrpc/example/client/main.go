package main

import (
	pb "awesomeProject/ES/hellogrpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const Address  = "127.0.0.1:50052"
const OpenTLS  = true

type customCredential struct {
}

func (c customCredential)GetRequestMetadata(ctx context.Context,uri ...string)(map[string]string,error)  {
	return map[string]string{
		"appid": "101010",
		"appkey": "i am key",
	},nil
}
func (c customCredential)RequireTransportSecurity() bool  {
	if OpenTLS{
		return true
	} else {
		return false
	}

}
func main()  {
	var opts []grpc.DialOption
	if OpenTLS{
		creds, err := credentials.NewClientTLSFromFile("./ES/hellogrpc/keys/server.pem", "TXXY")
		if err != nil{
			log.Fatalf("Failed to create tls credentials %v",err)
		}
		opts = append(opts,grpc.WithTransportCredentials(creds))
	}else {
		opts = append(opts,grpc.WithInsecure())
	}
	opts = append(opts,grpc.WithPerRPCCredentials(new(customCredential)))
	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c :=pb.NewHelloClient(conn)

	//调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name="gRPC"

	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r.Message)
}
