package main

import (
	"fmt"
	"image"
	"os"
	"testing"
)

var complianceSingleton *ComplianceHandler

func getSingleton() *ComplianceHandler {
	if complianceSingleton == nil {
		if complianceSingleton == nil {
			fmt.Println("initializing compliance handler")
			c := NewComplianceHandler()
			complianceSingleton = &c
		} else {
			fmt.Println("using initialized compliance handler")
		}
	} else {
		fmt.Println("using initialized compliance handler.")
	}

	return complianceSingleton
}
func loadAsImage(fileName string) image.Image {
	imgFile, _ := os.Open(fileName)
	img, _, _ := image.Decode(imgFile)
	return img
}

func TestContainsDog(t *testing.T) {
	return
	// given an image of a dog, and a person
	c := getSingleton()
	dog := loadAsImage("test_images/dog.jpg")
	person := loadAsImage("test_images/person.jpg")
	// when the computer vision examines both images
	println("checking for dogs in a picture of a dog")
	dogIsDog := c.containsDog(dog)
	println("checking for dogs in a picture without dogs")
	personIsDog := c.containsDog(person)
	// then the result will be that a dog is a dog, and a person is not.
	if !dogIsDog {
		println("dog is not dog, test failed")
		t.Fail()
	}
	if personIsDog {
		println("person is dog, test failed")
		t.Fail()
	}
}

func TestContainsHuman(t *testing.T) {
	// given an image of a dog, and a person
	c := getSingleton()
	dog := loadAsImage("test_images/dog.jpg")
	person := loadAsImage("test_images/person.jpg")
	// when the computer vision examines both images
	println("checking for human in a picture of a human")
	personIsPerson := c.containsHuman(person)
	println("checking for human in a picture of a dog")
	dogIsPerson := c.containsHuman(dog)
	// then the result will be that a person is a person, and a dog is not.
	if !personIsPerson {
		println("person is not person, test failed")
		t.Fail()
	}
	if dogIsPerson {
		println("dog is person, test failed")
		t.Fail()
	}
}

func TestIsCompliantDogAndPerson(t *testing.T) {
	return
	// given an image of a dog and a person
	c := getSingleton()
	dogAndPerson := loadAsImage("test_images/dog_and_person.jpg")
	// when the computer vision examines the image
	imgIsCompliant := c.IsCompliant(dogAndPerson)
	// then the result will be that the image is not compliant, as there is a person.
	if imgIsCompliant {
		t.Fail()
	}
}
