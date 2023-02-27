package main

import (
	"image"
	"os"
	"testing"
)

func loadAsImage(fileName string) image.Image {
	imgFile, _ := os.Open(fileName)
	img, _, _ := image.Decode(imgFile)
	return img
}

func TestContainsDog(t *testing.T) {
	// given an image of a dog, and a person
	dog := loadAsImage("test_images/dog.jpg")
	person := loadAsImage("test_images/person.jpg")
	// when the computer vision examines both images
	dogIsDog := containsDog(dog)
	personIsDog := containsDog(person)
	// then the result will be that a dog is a dog, and a person is not.
	if !dogIsDog || personIsDog {
		t.Fail()
	}
}

func TestContainsHuman(t *testing.T) {
	// given an image of a dog, and a person
	dog := loadAsImage("test_images/dog.jpg")
	person := loadAsImage("test_images/person.jpg")
	// when the computer vision examines both images
	personIsPerson := containsHuman(person)
	dogIsPerson := containsHuman(dog)
	// then the result will be that a person is a person, and a dog is not.
	if !personIsPerson || dogIsPerson {
		t.Fail()
	}
}

func TestIsCompliantDogAndPerson(t *testing.T) {
	// given an image of a dog and a person
	dogAndPerson := loadAsImage("test_images/dog_and_person.jpg")
	// when the computer vision examines the image
	imgIsCompliant := isCompliant(dogAndPerson)
	// then the result will be that the image is not compliant, as there is a person.
	if imgIsCompliant {
		t.Fail()
	}
}
