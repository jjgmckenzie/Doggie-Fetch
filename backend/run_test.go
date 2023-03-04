package main

import "testing"

func TestRunnerErrorsIfNoNetwork(t *testing.T) {
	_, err := NewRunner("", "", "", 0, 0, "", "")
	if err == nil {
		t.Fail()
	}
}

func TestRunnerErrorsIfNoAppId(t *testing.T) {
	_, err := NewRunner("./yolov3/yolov3.weights", "./yolov3/yolov3.cfg", "./yolov3/coco.names", 0, 0, "", "")
	if err == nil {
		t.Fail()
	}
}
