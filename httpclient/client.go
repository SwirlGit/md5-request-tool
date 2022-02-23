package httpclient

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
	keepAlive      = 30 * time.Second
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Config struct {
	Timeout             time.Duration
	IdleConnTimeout     time.Duration
	TLSHandshakeTimeout time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
}

type Client struct {
	httpClient httpClient
}

func NewClient(cfg Config) *Client {
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}

	client := http.Client{
		Transport: &http.Transport{
			IdleConnTimeout:     cfg.IdleConnTimeout,
			TLSHandshakeTimeout: cfg.TLSHandshakeTimeout,
			MaxIdleConns:        cfg.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
			}).DialContext,
			ForceAttemptHTTP2: true,
		},
		Timeout: timeout,
	}
	return &Client{httpClient: &client}
}

func (c *Client) GetResponseBody(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	return io.ReadAll(resp.Body)
}
