package cicd

import (
	"fmt"
	"os/exec"
	"strings"
)

type commands [][]string

func processor(job Job) error {

	err := build(job.PreBuild)
	if err != nil {
		return err
	}

	err = build(job.PreBuild)
	if err != nil {
		return err
	}

	err = build(job.PreBuild)
	if err != nil {
		return err
	}

	return nil
}

func build(statements []string) error {

	for _, statement := range statements {
		command := makeCommands(statement)
		args := command[1:]
		output := exec.Command(command[0], args...).Run()
		fmt.Println(output)
	}

	return nil
}

func makeCommands(statement string) []string {
	commands := strings.Split(statement, " ")
	return commands
}
