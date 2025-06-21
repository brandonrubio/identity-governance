package main

import (
	"github.com/brandonrubio/identity-governance/container"
)

func main() {
	container := container.NewContainer()
	container.Provide()
	container.Invoke()
}
