package di

import (
	"os"

	"go.uber.org/dig"

	"github.com/a3d21/goddd/application"
	"github.com/a3d21/goddd/domain/account"
	"github.com/a3d21/goddd/infra/meminfra"
	"github.com/a3d21/goddd/infra/prodinfra"
)

//go:generate go test . -gen
func GetContainerByEnv() *dig.Container {
	switch os.Getenv("APP_ENV") {
	case "mem":
		return newMemContainer()
	case "prod":
		return newProdContainer()
	case "test":
		return newTestContainer()
	default:
		return newMemContainer()
	}
}

// newMemContainer 内存环境DI容器，用于测试Domain
func newMemContainer() *dig.Container {
	c := dig.New()
	c.Provide(application.NewHttpApp)
	c.Provide(account.NewService)
	c.Provide(meminfra.NewAccountRepo)
	return c
}

// newProdContainer 生产环境DI容器，使用真实DB、MQ连接
func newProdContainer() *dig.Container {
	c := dig.New()
	c.Provide(application.NewHttpApp)
	c.Provide(account.NewService)
	c.Provide(prodinfra.NewAccountRepo)
	c.Provide(prodinfra.NewDB)
	return c
}

// newTestContainer 测试环境DI容器，使用sqlite fake db
func newTestContainer() *dig.Container {
	c := dig.New()
	c.Provide(application.NewHttpApp)
	c.Provide(account.NewService)
	c.Provide(prodinfra.NewAccountRepo)
	c.Provide(prodinfra.NewStubDB)
	return c
}
