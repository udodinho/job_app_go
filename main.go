package main

import (
	// "github.com/udodinho/job-app/app/controllers"
	"github.com/udodinho/job-app/domain/entity"
	// "github.com/udodinho/job-app/app/server"
	// "github.com/udodinho/job-app/application/server"
	"github.com/udodinho/job-app/app/server"
	// "github.com/udodinho/job-app/infrastructure/repository"
)

func main() {

	db := entity.DB
	// db := repository.DB{}

	// app := &controllers.Application{
	// 	Db: &db,
	// }

	app := server.Server{
		Db: db,
	}

	s := server.Server(app)
	// s := server.NewServer(app)
	s.Start()
}