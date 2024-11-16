package main

import (
	"fmt"
	"funtwitt/config"
	user_controller "funtwitt/domain/user/controller"
	user_repository "funtwitt/domain/user/repository"
	user_service "funtwitt/domain/user/service"
	"funtwitt/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start Server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	if err != nil {
		panic(err)
	}

	// Graceful Shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		log.Println("Received shutdown signal, closing database connection.")
		db.Prisma.Disconnect()
		os.Exit(0)
	}()

	userRepository := user_repository.NewUserRepository(db)
	userService := user_service.NewUserServiceImpl(userRepository)
	userController := user_controller.NewController(userService)

	routes := router.NewRouter(
		userController,
	)

	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	allowedMethods := strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")
	allowedHeaders := strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ",")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedMethods(allowedMethods),
		handlers.AllowedHeaders(allowedHeaders),
	)(routes)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        corsHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server running on port %s\n", port)
	server_err := server.ListenAndServe()

	if server_err != nil {
		panic(server_err)
	}
}
