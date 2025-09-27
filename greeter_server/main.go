package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/IdentityStolen/HelloGRPC/helloworld"
	"google.golang.org/grpc"
)

type myHelloServer struct {
	pb.UnimplementedGreeterServer
}

func (s *myHelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received %v", req.GetName())
	return &pb.HelloReply{Message: fmt.Sprintf("Hi %v, welcome to the wizarding world!", req.GetName())}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("Error creating listener: %q", err)
	}
	serverRegister := grpc.NewServer()
	service := &myHelloServer{}
	pb.RegisterGreeterServer(serverRegister, service)
	log.Printf("starting server and serving requests...")
	err = serverRegister.Serve(listener)
	if err != nil {
		log.Fatalf("Something went wrong: %q", err)
	}
}
