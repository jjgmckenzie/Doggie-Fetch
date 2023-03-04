package main

import "log"

func main() {
	log.Fatalln(DefaultRunner(":8080").run("dog-api-images"))
}
