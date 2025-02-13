package main

import (
	"io"
	"net/http"
)

func GetPokeLocations(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
