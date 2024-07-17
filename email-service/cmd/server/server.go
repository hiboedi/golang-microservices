package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "email-service/pb"
	"email-service/pkg/email"
)

type server struct {
	pb.UnimplementedEmailServiceServer
}

func (s *server) SendEmail(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	log.Printf("Request from client: %v\n", req)
	err := email.SendEmail(req.To, req.PlayerId, req.ProductName, req.ProductPrice)
	if err != nil {
		return &pb.EmailResponse{Status: "FAILED"}, err
	}
	return &pb.EmailResponse{Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEmailServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Server running at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
