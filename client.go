package dagobert

import (
	"net/url"

	jina "github.com/jina-ai/client-go"
)

type Client struct {
	scheme string
	host   string
	client jina.Client
}

func NewClient(uri string) (*Client, error) {
	clipURL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	c := Client{
		scheme: clipURL.Scheme,
		host:   clipURL.Host,
	}

	switch c.scheme {
	case "", "grpc", "grpcs":
		c.client, err = jina.NewGRPCClient(c.host)
	case "ws", "wss", "websocket":
		c.client, err = jina.NewWebSocketClient(c.host)
	case "http", "https":
		c.client, err = jina.NewHTTPClient(c.host)
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}
