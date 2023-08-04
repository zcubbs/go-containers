package gocontainers

import (
	"testing"
	"time"
)

func TestRunContainerOnHost(t *testing.T) {
	err := RunContainerOnHost("docker.io/library/alpine", []string{"echo", "hello world"})
	if err != nil {
		t.Error(err)
	}

	time.Sleep(5 * time.Second)
}
