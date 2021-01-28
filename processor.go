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

	err = build(job.Build)
	if err != nil {
		return err
	}

	err = build(job.PostBuild)
	if err != nil {
		return err
	}

	return nil
}

func build(statements []string) error {

	for _, statement := range statements {
		command := makeCommands(statement)
		var args []string
		if len(command) > 1 {
			args = command[1:]
		}

		output, err := exec.Command(command[0], args...).Output()
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	}

	return nil
}

func makeCommands(statement string) []string {
	commands := strings.Split(statement, " ")
	return commands
}
