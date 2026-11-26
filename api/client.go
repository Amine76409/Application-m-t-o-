package api

import (
	"net/http"
	"time"
)

var Client = &http.Client{Timeout: 5 * time.Second}

func Get(adresse string) (*http.Response, error) {
	req, err := http.NewRequest("GET", adresse, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "tp-meteo-ynov/1.0")

	return Client.Do(req)
}
