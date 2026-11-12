package main

import (
	"Di-Test/internal/config"
	"Di-Test/internal/database"
	"Di-Test/internal/handler"
	"Di-Test/internal/repository"
	"Di-Test/internal/router"
	"Di-Test/internal/service"
	"fmt"
	"net/http"
)

func main() {
	cfg := config.NewConfig()
	if cfg == nil {
		panic("Failed to load configuration")
	}
	db := database.NewMongoDB(cfg.MongoURI, cfg.DBName)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	r := router.Router(userHandler)

	// Use cfg.Port as needed

	fmt.Println("Server is run ", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, r)

}
