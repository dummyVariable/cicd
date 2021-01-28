package main

import (
	"io/ioutil"
	"log"
)

func main() {

	file, err := ioutil.ReadFile("test2.yaml")
	if err != nil {
		log.Fatal(err)
	}

	buildJobs := jobs(file)

	err = processor(buildJobs)
	if err != nil {
		log.Fatal(err)
	}

}
