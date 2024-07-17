package handler

import (
	"context"
	"log"

	pb "email-service/pb"
	"email-service/pkg/email"
)

type EmailService struct {
	pb.UnimplementedEmailServiceServer
}

func (s *EmailService) SendEmail(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	log.Printf("Request from client: %v\n", req)
	err := email.SendEmail(req.To, req.PlayerId, req.ProductName, req.ProductPrice)
	if err != nil {
		return &pb.EmailResponse{Status: "FAILED"}, err
	}
	return &pb.EmailResponse{Status: "OK"}, nil
}
