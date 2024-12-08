package youtube

import (
	"net/http"
	"time"

	yt2 "github.com/kkdai/youtube/v2"
)

type Client struct {
	HttpClient   *http.Client
	Timeout      time.Duration
	streamClient yt2.Client
}

func New(cli *http.Client) Client {
	if cli == nil {
		cli = http.DefaultClient
	}

	return Client{
		HttpClient: cli,
	}
}
