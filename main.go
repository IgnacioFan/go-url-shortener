package main

import "go-url-shortener/cmd"

func main() {
  cmd.Execute()
}

// type Server struct {
//   rpc_api.ShortenerServer
// 	Url url_service.UrlService
// }

// func (s *Server) ShortenURL(ctx context.Context, req *rpc_api.URL) (*rpc_api.ShortenedURL, error) {
//   if !govalidator.IsURL(req.LongUrl) {
//     return nil, fmt.Errorf("invalid url")
//   }
// 	shortURL, err := s.Url.GenerateShortURL(req.LongUrl)
//   if err != nil {
//     return nil, err
//   }
// 	return &rpc_api.ShortenedURL{ShortUrl: shortURL, LongUrl: req.LongUrl}, nil
// }

// func (s *Server) RedirectURL(ctx context.Context, req *rpc_api.ShortenedURL) (*rpc_api.URL, error) {
//   if !govalidator.Matches(req.ShortUrl, SHROT_URL_REGEX) {
//     return nil, fmt.Errorf("invalid url")
//   }
//   longURL, err := s.Url.OriginalURL(req.ShortUrl)
//   if err != nil {
//     return nil, err
//   }
// 	return &rpc_api.URL{LongUrl: longURL}, nil
// }

// func init()  {
// 	godotenv.Load()
// }

// func main() {
// 	fmt.Println("starting gRPC server...")
// 	db, err := postgres.NewPostgres()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1)
// 	}

// 	if err = db.NewMigrate(); err != nil {
// 		fmt.Println(err.Error())
// 	  os.Exit(1)
// 	}

//  service, err := url_service.NewUrlService()
// 	if err != nil {
// 	  fmt.Println(err.Error())
// 	  os.Exit(1)
// 	}

//  lis, err := net.Listen("tcp", "localhost:50051")
//  if err != nil {
//   log.Fatalf("failed to listen: %v \n", err)
//  }

//  grpcServer := grpc.NewServer()
//  rpc_api.RegisterShortenerServer(grpcServer, &rpc_api.Server{Url: service})

//  if err := grpcServer.Serve(lis); err != nil {
//   log.Fatalf("failed to serve: %v \n", err)
//  }
// }
