package piped

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type instance struct {
	Name                 string  `json:"name"`
	ApiURL               string  `json:"api_url"`
	Locations            string  `json:"locations"`
	Version              string  `json:"version"`
	UpToDate             bool    `json:"up_to_date"`
	Cdn                  bool    `json:"cdn"`
	Registered           int     `json:"registered"`
	LastChecked          int64   `json:"last_checked"`
	Cache                bool    `json:"cache"`
	S3Enabled            bool    `json:"s3_enabled"`
	ImageProxyURL        string  `json:"image_proxy_url"`
	RegistrationDisabled bool    `json:"registration_disabled"`
	Uptime24h            float64 `json:"uptime_24h"`
	Uptime7d             float64 `json:"uptime_7d"`
	Uptime30d            float64 `json:"uptime_30d"`
}

type PipedClient struct {
	instances []instance
}

func New() (PipedClient, error) {
	p := PipedClient{}

	instances, err := p.fetchInstances()
	if err != nil {
		return p, err
	}

	p.instances = instances

	return p, nil
}

func (p PipedClient) fetchInstances() ([]instance, error) {
	url := "https://piped-instances.kavin.rocks/"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var instances []instance

	if err := json.Unmarshal(body, &instances); err != nil {
		return nil, err
	}

	return instances, nil
}
