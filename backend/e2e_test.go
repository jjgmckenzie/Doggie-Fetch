package main

import (
	"bytes"
	"encoding/base64"
	"github.com/goccy/go-json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

type mockImagePost struct {
	Name  string
	Breed string
	Image string
}

func TestEndToEnd(t *testing.T) {
	go func() {
		_ = run("test")
	}()
	time.Sleep(2 * time.Second) // wait for Gin to Initialize.
	loadImage, _ := os.ReadFile("./test_images/large_dog_image.jpg")
	imageText := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(loadImage)
	mock := mockImagePost{
		Name:  "Big Dog Image",
		Breed: "mix",
		Image: imageText,
	}
	jsonMock, _ := json.Marshal(mock)
	resp, err := http.Post("http://localhost:8080/upload", "application/json", bytes.NewReader(jsonMock))
	if err != nil {
		println(err.Error())
		t.Fail()
	}
	if resp.StatusCode != http.StatusAccepted {
		println(resp.StatusCode)
		t.Fail()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(body), "https://github.com/GoFetchBot/test/pull/") {
		println("received wrong format: " + string(body))
		t.Fail()
	}
}
