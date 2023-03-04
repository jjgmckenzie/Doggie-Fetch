package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestMainOpensServer(t *testing.T) {
	go main()
	time.Sleep(2 * time.Second) // wait for Gin to Initialize.
	mock := mockImagePost{
		Name:  "Big Dog Image",
		Breed: "mix",
		Image: "not_an_image",
	}
	jsonMock, _ := json.Marshal(mock)
	resp, err := http.Post("http://localhost:8000/upload", "application/json", bytes.NewReader(jsonMock))
	if err != nil {
		println(err.Error())
		t.Fail()
	}
	if resp.StatusCode != http.StatusBadRequest {
		println(resp.StatusCode)
		t.Fail()
	}
}
