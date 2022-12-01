package main

import (
	"SFA/week11/Lecture30/grpc/pb"
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":9000"

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost"+port, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)
	ctx := context.Background()
	req := &pb.TopCountRequest{MaxCount: 10}
	resp, err := client.FetchTopStories(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	stories, err := client.FetchItems(ctx, resp)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range stories.TopStories {
		fmt.Printf("[%d]: %v\n", i+1, v)
	}
}
