package postedimage

import (
	"log"
	"testing"
)

func TestProfanityFilter(t *testing.T) {
	// given a profane word and a not profane, similar word
	profaneWord := "ass"
	notProfaneWord := "lassie"
	// when the profanity filter considers them
	_, profaneWordProfane := filter(profaneWord)
	_, notProfaneWordProfane := filter(notProfaneWord)
	// then the profane word will error, and the not profane word will not
	if profaneWordProfane == nil {
		t.Fail()
	}
	if notProfaneWordProfane != nil {
		t.Fail()
	}
}

func TestBuilderBuild(t *testing.T) {
	img, err := New().Build("maya", "pyrenees", encodeTestImage("jpg"))
	if err != nil {
		log.Printf("an error occured: %s", err.Error())
		t.Fail()
		return
	}
	r, g, b, a := img.Image.At(0, 0).RGBA()
	if r>>8 != 192 || g>>8 != 192 || b>>8 != 192 || a>>8 != 255 {
		t.Fail()
	}
}

func TestTruncateString(t *testing.T) {
	if truncate("test") != "test" {
		t.Fail()
	}
	if len(truncate("abcdefghijklmnopqrstuvwxyz")) != 20 {
		t.Fail()
	}
}
