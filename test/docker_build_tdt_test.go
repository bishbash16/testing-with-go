package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerHelloWorldTDT(t *testing.T) {
	// Configure the tag to use on the Docker image.
	tag := "hello-world:go-test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	// Build the Docker image.
	docker.Build(t, "../", buildOptions)

	// A testing table to test different aspects of the image
	tt := []struct {
		name           string
		command        []string
		expectedOutput string
		errorMessage   string
	}{
		{
			name:           "Test case 1: Check if the file exists",
			command:        []string{"sh", "-c", "if [ -f /test.txt ]; then echo 'File exists'; else echo 'No such file exists'; fi"},
			expectedOutput: "File exists",
			errorMessage:   "File does not exist",
		},
		{
			name:           "Test case 2: Check file contents",
			command:        []string{"sh", "-c", "cat /test.txt"},
			expectedOutput: "Hello, World!",
			errorMessage:   "File contents do not match",
		},
	}

	// Iterate over the test table and run the tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Allow tests to run in parallel
			t.Parallel()

			// The docker run options
			opts := &docker.RunOptions{
				// Remove the container once finished
				Remove: true,

				// The command we will run for the test
				Command: tc.command,
			}

			// Run the container, and get the output
			output := docker.Run(t, tag, opts)

			// The test check to assert we get what we expected.
			assert.Equal(t, tc.expectedOutput, output, tc.errorMessage)
		})
	}
}
