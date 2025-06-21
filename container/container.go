package container

import (
	"github.com/brandonrubio/identity-governance/service/logger"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

type Dependencies struct {
	dig.In
	LoggerService logger.ILoggerService `name:"LoggerService"`
}

func (c *Container) Provide() {
	c.container.Provide(logger.NewLoggerService, dig.As(new(logger.ILoggerService)), dig.Name("LoggerService"))
}

func (c *Container) Invoke() {
	c.container.Invoke(func(deps Dependencies) {
		deps.LoggerService.Init()
		deps.LoggerService.Log("`Hello, World!")
	})
}

func NewContainer() *Container {
	return &Container{
		container: dig.New(),
	}
}
