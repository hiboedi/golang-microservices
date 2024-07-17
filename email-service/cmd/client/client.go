package main

import (
	"context"
	"log"
	"time"

	pb "email-service/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEmailServiceClient(conn)

	email := &pb.EmailRequest{
		To:           "hi.boedi8@gmail.com",
		PlayerId:     "123456",
		ProductName:  "PUBG",
		ProductPrice: 100000,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SendEmail(ctx, email)
	if err != nil {
		log.Fatalf("could not send email: %v", err)
	}
	log.Printf("Response from server: %v", resp.Status)
}
