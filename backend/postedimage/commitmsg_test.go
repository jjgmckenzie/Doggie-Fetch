package postedimage

import "testing"

func TestBreedFormatter(t *testing.T) {
	// given unformatted breeds
	borderCollie := "collie-border"
	pyrenees := "pyrenees"
	yorkshireTerrier := "terrier-yorkshire"
	// when formatBreed() & capitalize() is called on them, then they are formatted correctly
	if capitalize(formatBreed(borderCollie)) != "Border Collie" {
		t.Fail()
	}
	if capitalize(formatBreed(pyrenees)) != "Pyrenees" {
		t.Fail()
	}
	if capitalize(formatBreed(yorkshireTerrier)) != "Yorkshire Terrier" {
		t.Fail()
	}
}

func TestCommitMessage(t *testing.T) {
	// given an image
	img := Image{
		Name:  "MAYA",
		Breed: "pyrenees",
	}
	// when the commit message is requested
	actual := img.CommitMessage()
	expected := "Adds a Pyrenees named Maya"
	// then it matches the expected commit message
	if actual != expected {
		t.Fail()
	}
}
