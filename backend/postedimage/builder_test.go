package postedimage

import (
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

func TestTruncateString(t *testing.T) {
	if truncate("test") != "test" {
		t.Fail()
	}
	if len(truncate("abcdefghijklmnopqrstuvwxyz")) != 20 {
		t.Fail()
	}
}
