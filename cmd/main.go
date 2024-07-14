package main

import (
	"fmt"

	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Should put this into config file or GSM
const (
	dbHost     = "db"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "database"
)

func main() {
	e := echo.New()

	// var server generated.ServerInterface = newServer()

	// generated.RegisterHandlers(e, server)

	var server handler.ServerInterface = newServer()

	e.POST("/estate", server.AddEstate)
	e.POST("/estate/:id/tree", server.AddTree)
	e.GET("/estate/:id/stats", server.GetEstateStats)
	e.GET("/estate/:id/droneplan", server.GetDronePlan)

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() handler.ServerInterface {
	dbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	var service service.ServiceInterface = service.NewService(repo)
	opts := handler.NewServerOptions{
		Service: service,
	}
	return handler.NewServer(opts)
}
