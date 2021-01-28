package cicd

import (
	"log"

	"gopkg.in/yaml.v2"
)

//Job defines a structure of job instuctions
type Job struct {
	PreBuild  []string `yaml:"prebuild"`
	Build     []string `yaml:"build"`
	PostBuild []string `yaml:"postbuild"`
}

func jobs(file []byte) Job {

	var jobs Job

	err := yaml.Unmarshal(file, &jobs)
	if err != nil {
		log.Fatal(err)
	}

	return jobs
}
