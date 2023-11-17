package main

import (
	"github.com/go-playground/validator/v10"
	"go-gin-goorm/config"
	"go-gin-goorm/controller"
	"go-gin-goorm/helper"
	"go-gin-goorm/models"
	"go-gin-goorm/repository"
	"go-gin-goorm/router"
	"go-gin-goorm/service"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server started")
	//Database
	db := config.DatabaseConnection()

	validate := validator.New()

	db.Take("tags").AutoMigrate(&models.Tags{})

	// Repository

	tagRepository := repository.NewTagsRepositoryImpl(db)

	//Service
	tagsService := service.NewTagsServiceImpl(tagRepository, validate)

	// Controller
	tagsController := controller.NewsTagController(tagsService)
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()

	helper.ErrorPanic(err)

}
