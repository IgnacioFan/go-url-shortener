package grpccontroller

import (
	"context"
	"fmt"
	"log"
	"net"

	urlshortnerPB "go-url-shortener/url_shortner_service/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (*Server) GetShortUrl(ctx context.Context, req *urlshortnerPB.ShrotUrlRequest) (*urlshortnerPB.ShortUrlResponse, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	url := req.GetLongUrl()
	fmt.Println(url)

	res := &urlshortnerPB.ShortUrlResponse{
		ShortUrl: "",
	}

	return res, nil
}

func (*Server) GetLongUrl(ctx context.Context, req *urlshortnerPB.LongUrlRequest) (*urlshortnerPB.LongUrlResponse, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	shortUrl := req.GetShortUrl()
	fmt.Println(shortUrl)

	res := &urlshortnerPB.LongUrlResponse{
		LongUrl: "",
	}

	return res, nil
}

func NewServer() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	urlshortnerPB.RegisterUrlShortnerServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
