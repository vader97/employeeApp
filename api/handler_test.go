package api

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"employeeApp/models/employee"
	repository "employeeApp/repository/employee"
	service "employeeApp/service/employee"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestEmployeeHandler(t *testing.T) {
	router := setupRouter()

	repo := repository.NewInMemoryEmployeeRepository()
	service := service.NewEmployeeService(repo)
	handler := NewEmployeeHandler(service)

	router.GET("/employee/:id", handler.GetEmployeeByID)
	router.POST("/employee", handler.CreateEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)
	router.GET("/employees", handler.ListEmployees)

	// Test GetEmployeeByID
	emp := &employee.Employee{ID: 1, Name: "John Doe", Position: "Developer", Salary: 50000}
	repo.CreateEmployee(emp)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employee/1", nil)
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Test CreateEmployee
	jsonEmp := `{"id":2,"name":"Jane Smith","position":"Manager","salary":60000}`
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/employee", strings.NewReader(jsonEmp))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Test UpdateEmployee
	//updatedEmp := &employee.Employee{ID: 2, Name: "Jane Smith", Position: "Manager", Salary: 70000}
	jsonUpdatedEmp := `{"id":2,"name":"Jane Smith","position":"Manager","salary":70000}`
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/employee/2", strings.NewReader(jsonUpdatedEmp))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Test DeleteEmployee
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/employee/2", nil)
	router.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
	}

	// Test ListEmployees
	for i := 0; i < 15; i++ {
		emp := &employee.Employee{ID: i + 1, Name: "Employee" + strconv.Itoa(i+1), Position: "Position" + strconv.Itoa(i+1), Salary: float64((i + 1) * 1000)}
		repo.CreateEmployee(emp)
	}
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/employees?page=1&size=10", nil)
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
