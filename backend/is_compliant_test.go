package main

import (
	"fmt"
	"image"
	"os"
	"sync"
	"testing"
)

var complianceSingleton *ComplianceHandler
var lock = &sync.Mutex{}

func getSingleton() *ComplianceHandler {
	if complianceSingleton == nil {
		lock.Lock()
		defer lock.Unlock()
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
	// given an image of a dog, and a person
	c := getSingleton()
	dog := loadAsImage("test_images/dog.jpg")
	person := loadAsImage("test_images/person.jpg")
	dogContents, _ := c.getContents(dog)
	personContents, _ := c.getContents(person)
	// when the computer vision examines both images
	println("checking for dogs in a picture of a dog")
	dogIsDog := c.containsDog(dogContents)
	println("checking for dogs in a picture without dogs")
	personIsDog := c.containsDog(personContents)
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
	dogContents, _ := c.getContents(dog)
	personContents, _ := c.getContents(person)
	// when the computer vision examines both images
	println("checking for human in a picture of a human")
	personIsPerson := c.containsHuman(personContents)
	println("checking for human in a picture of a dog")
	dogIsPerson := c.containsHuman(dogContents)
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
	// given an image of a dog and a person
	c := getSingleton()
	dogAndPerson := loadAsImage("test_images/dog_and_person.jpg")
	// when the computer vision examines the image
	imgIsCompliant, _ := c.IsCompliant(dogAndPerson)
	// then the result will be that the image is not compliant, as there is a person.
	if imgIsCompliant {
		t.Fail()
	}
}
