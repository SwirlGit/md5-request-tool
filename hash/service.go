package hash

import (
	"context"
	"crypto/md5"
)

type responseGetter interface {
	GetResponseBody(ctx context.Context, url string) ([]byte, error)
}

type Service struct {
	responseGetter responseGetter
}

func NewService(responseGetter responseGetter) *Service {
	return &Service{responseGetter: responseGetter}
}

func (s *Service) MD5(ctx context.Context, url string) ([md5.Size]byte, error) {
	resp, err := s.responseGetter.GetResponseBody(ctx, url)
	if err != nil {
		return [md5.Size]byte{}, err
	}

	return md5.Sum(resp), nil
}