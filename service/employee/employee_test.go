package employee

import (
	"strconv"
	"testing"

	"employeeApp/models/employee"
	repository "employeeApp/repository/employee"
)

func TestEmployeeService(t *testing.T) {
	repo := repository.NewInMemoryEmployeeRepository()
	service := NewEmployeeService(repo)

	// Test CreateEmployee
	emp := &employee.Employee{ID: 1, Name: "John Doe", Position: "Developer", Salary: 50000}
	if err := service.CreateEmployee(emp); err != nil {
		t.Errorf("Error creating employee: %v", err)
	}

	// Test GetEmployeeByID
	retrievedEmp, err := service.GetEmployeeByID(1)
	if err != nil {
		t.Errorf("Error getting employee by ID: %v", err)
	}
	if retrievedEmp == nil || retrievedEmp.ID != 1 {
		t.Errorf("Invalid employee retrieved")
	}

	// Test UpdateEmployee
	emp.Salary = 60000
	if err := service.UpdateEmployee(emp); err != nil {
		t.Errorf("Error updating employee: %v", err)
	}
	retrievedEmp, _ = service.GetEmployeeByID(1)
	if retrievedEmp.Salary != 60000 {
		t.Errorf("Employee update failed")
	}

	// Test DeleteEmployee
	if err := service.DeleteEmployee(1); err != nil {
		t.Errorf("Error deleting employee: %v", err)
	}
	retrievedEmp, err = service.GetEmployeeByID(1)
	if retrievedEmp != nil || err == nil {
		t.Errorf("Employee deletion failed")
	}

	// Test ListEmployees
	for i := 0; i < 15; i++ {
		emp := &employee.Employee{ID: i + 1, Name: "Employee" + strconv.Itoa(i+1), Position: "Position" + strconv.Itoa(i+1), Salary: float64((i + 1) * 1000)}
		repo.CreateEmployee(emp)
	}
	employees, err := service.ListEmployees(1, 10)
	if err != nil {
		t.Errorf("Error listing employees: %v", err)
	}
	if len(employees) != 10 {
		t.Errorf("Expected 10 employees, got %d", len(employees))
	}
	if employees[0].ID != 1 || employees[9].ID != 10 {
		t.Errorf("Incorrect employee pagination")
	}
}
