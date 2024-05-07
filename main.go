package main

import (
	"employeeApp/api"
	employeeRepo "employeeApp/repository/employee"
	employeeService "employeeApp/service/employee"
	"github.com/gin-gonic/gin"
)

func main() {
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

	// Start server
	r.Run(":80")
}
