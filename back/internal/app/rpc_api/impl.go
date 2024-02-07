package rpc_api

import (
	"context"
	"fmt"
	"go-url-shortener/internal/service/url_service"

	"github.com/asaskevich/govalidator"
)

const (
  SHROT_URL_REGEX = "^[a-zA-Z0-9]{1,7}$"
)

type Server struct {
  ShortenerServer
	Url url_service.UrlService
}

func (s *Server) ShortenURL(ctx context.Context, req *URL) (*ShortenedURL, error) {
  if !govalidator.IsURL(req.LongUrl) {
    return nil, fmt.Errorf("invalid url")
  }
	shortURL, err := s.Url.GenerateShortURL(req.LongUrl)
  if err != nil {
    return nil, err
  }
	return &ShortenedURL{ShortUrl: shortURL, LongUrl: req.LongUrl}, nil
}

func (s *Server) RedirectURL(ctx context.Context, req *ShortenedURL) (*URL, error) {
  if !govalidator.Matches(req.ShortUrl, SHROT_URL_REGEX) {
    return nil, fmt.Errorf("invalid url")
  }
  longURL, err := s.Url.OriginalURL(req.ShortUrl)
  if err != nil {
    return nil, err
  }
	return &URL{LongUrl: longURL}, nil
}
