package controller

import (
	"context"
	"fmt"
	"log"
	"net"

	urlshortnerPB "go-url-shortener/url_shortner_service/proto"
	"go-url-shortener/url_shortner_service/service"

	"google.golang.org/grpc"
)

type UrlShortenerHandler struct {
	Service service.UrlShortenService
}

func (h *UrlShortenerHandler) GetShortUrl(ctx context.Context, req *urlshortnerPB.ShrotUrlRequest) (*urlshortnerPB.ShortUrlResponse, error) {
	fmt.Printf("GetShortUrl function is invoked with %v \n", req)

	url := req.GetLongUrl()
	res, err := h.Service.CreateShortUrl(url)
	if err != nil {
		fmt.Println(err)
	}
	return &urlshortnerPB.ShortUrlResponse{
		ShortUrl: res,
	}, nil
}

func (h *UrlShortenerHandler) GetLongUrl(ctx context.Context, req *urlshortnerPB.LongUrlRequest) (*urlshortnerPB.LongUrlResponse, error) {
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
	urlshortnerPB.RegisterUrlShortnerServiceServer(
		grpcServer,
		&UrlShortenerHandler{
			Service: *service.NewUrlShortenService(),
		},
	)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
