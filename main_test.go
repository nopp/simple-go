package main

import (
	"net/http"
	"simple-go/app"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	go app.StartApp()
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://localhost:8000/version", nil)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetRoot(t *testing.T) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://localhost:8000/", nil)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
