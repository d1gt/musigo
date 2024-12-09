package invidious

import (
	"encoding/json"
	"io"
	"net/http"
)

var invidiousInstancesUrl string = "https://api.invidious.io/instances.json?"

type instance struct {
	Name string
	Type string `json:"type"`
	URI  string `json:"uri"`
}
type Invidious struct {
	instances []instance
}

func New() Invidious {
	return Invidious{}
}

func (inv Invidious) fetchInstances() ([]instance, error) {
	resp, err := http.Get(invidiousInstancesUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var instancesJson [][]json.RawMessage
	var instances []instance

	if err := json.Unmarshal([]byte(body), &instancesJson); err != nil {
		return nil, err
	}

	for _, entry := range instancesJson {
		var instance instance

		if err := json.Unmarshal(entry[0], &instance.Name); err != nil {
			continue
		}

		if err := json.Unmarshal(entry[1], &instance); err != nil {
			continue

		}

		if instance.Type == "https" {
			instances = append(instances, instance)
		}

	}

	return instances, nil
}
