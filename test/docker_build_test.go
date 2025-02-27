package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerHelloWorld(t *testing.T) {
	// Configure the tag to use on the Docker image.
	tag := "hello-world:go-test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	// Build the Docker image.
	docker.Build(t, "../", buildOptions)

	// Test case 1: Run the Docker image, read the text file from it, and make sure it contains the expected output.
	opts := &docker.RunOptions{
		Remove: true,
		Command: []string{
			"cat",
			"/test.txt"}}

	output := docker.Run(t, tag, opts)
	assert.Equal(t, "Hello, World!", output)

	// Test case 2: Check if the file exists in the Docker container.
	opts = &docker.RunOptions{
		Remove: true,
		Command: []string{
			"/bin/sh",
			"-c",
			"if [ -f /test.txt ]; then echo 'File exists'; else echo 'No such file exists'; fi"}}
	output = docker.Run(t, tag, opts)
	assert.Equal(t, "File exists", output)
}
