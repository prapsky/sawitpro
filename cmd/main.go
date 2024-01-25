package main

import (
	"os"

	"github.com/prapsky/sawitpro/generated"
	"github.com/prapsky/sawitpro/handler"
	"github.com/prapsky/sawitpro/repository"
	"github.com/prapsky/sawitpro/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	repository := repository.NewRepository(os.Getenv("DATABASE_URL"))
	service := service.NewUserService(service.UserServiceOptions{
		Repository: repository,
		AuthService: service.NewJwtAuthService(service.JwtAuthServiceOptions{
			PrivateKey: os.Getenv("PRIVATE_KEY"),
			PublicKey:  os.Getenv("PUBLIC_KEY"),
		}),
	})

	return handler.NewServer(service)
}
