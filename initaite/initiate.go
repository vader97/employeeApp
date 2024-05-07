package initaite

import (
	"context"
	"employeeApp/api"
	employeeRepo "employeeApp/repository/employee"
	employeeService "employeeApp/service/employee"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitiateServer() {
	r := gin.Default()
	repo := employeeRepo.NewInMemoryEmployeeRepository()
	empService := employeeService.NewEmployeeService(repo)
	empHandler := api.NewEmployeeHandler(empService)
	empV1 := r.Group("/api/v1")
	{
		empV1.GET("/employee/:id", empHandler.GetEmployeeByID)
		empV1.POST("/employee", empHandler.CreateEmployee)
		empV1.PUT("/employee/:id", empHandler.UpdateEmployee)
		empV1.DELETE("/employee/:id", empHandler.DeleteEmployee)
		empV1.GET("/employees", empHandler.ListEmployees)
	}

	server := &http.Server{
		Addr:    ":80",
		Handler: r,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	// Wait for an interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Server is shutting down...")

	// Create a context with a timeout to allow outstanding requests to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown error: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}
