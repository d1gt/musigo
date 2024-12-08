package invidious

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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
	url := "https://api.invidious.io/instances.json?pretty=1&sort_by=type,users"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var instancesJson [][]json.RawMessage
	var instances []instance

	if err := json.Unmarshal([]byte(body), &instancesJson); err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}

	for _, entry := range instancesJson {
		var instance instance

		// Unmarshal each part separately
		if err := json.Unmarshal(entry[0], &instance.Name); err != nil {
			log.Printf("Failed to parse name: %v", err)
			continue
		}

		if err := json.Unmarshal(entry[1], &instance); err != nil {
			log.Printf("Failed to parse instance: %v", err)
			continue

		}

		if instance.Type == "https" {
			instances = append(instances, instance)
		}

	}

	return instances, nil
}
