# go-containers: Go Docker Containers package

[![tag](https://img.shields.io/github/tag/zcubbs/go-containers)](https://github.com/zcubbs/go-containers/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.20-%23007d9c)
[![GoDoc](https://godoc.org/github.com/zcubbs/go-containers?status.svg)](https://pkg.go.dev/github.com/zcubbs/go-containers)
![Build Status](https://github.com/zcubbs/go-containers/actions/workflows/test.yaml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/zcubbs/go-containers)](https://goreportcard.com/report/github.com/zcubbs/go-containers)
[![Coverage](https://img.shields.io/codecov/c/github/zcubbs/go-containers)](https://codecov.io/gh/zcubbs/go-containers)
[![Contributors](https://img.shields.io/github/contributors/zcubbs/go-containers)](https://github.com/zcubbs/go-containers/graphs/contributors)
[![License](https://img.shields.io/github/license/zcubbs/go-containers)](./LICENSE)

This Go package provides a simple interface for running Docker containers directly from a host system.

## Features

- **Pull Docker Image**: The `RunContainerOnHost` function pulls a Docker image from a Docker registry.
- **Create Docker Container**: The function creates a Docker container from the pulled image.
- **Run Docker Container**: It then runs the Docker container on the host.
- **Wait For Container To Complete**: The function waits for the Docker container to finish its execution.
- **Fetch Logs**: It fetches and prints the logs of the Docker container to `os.Stdout`.

## Dependencies

This package uses the following dependencies:

- `context`: Standard Go package for carrying deadlines, cancellations signals, and other request-scoped values across API boundaries.
- `github.com/docker/docker/api/types`: Docker API types.
- `github.com/docker/docker/api/types/container`: Docker container types.
- `github.com/docker/docker/client`: Docker client to interact with Docker API.
- `io`: Standard Go package for basic interfaces to I/O primitives.
- `os`: Standard Go package that provides a platform-independent interface to operating system functionality.

## Usage

```go
package main

import (
	"github.com/zcubbs/gocontainers"
)

func main() {
	image := "ubuntu:latest"
	cmd := []string{"/bin/sh", "-c", "echo Hello, World"}

	err := gocontainers.RunContainerOnHost(image, cmd)
	if err != nil {
		panic(err)
	}
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
