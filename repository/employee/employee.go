package employee

import (
	"employeeApp/models/customerrors"
	"employeeApp/models/employee"
	"sync"
)

// using mutex for locking the in memory db for concurrent requests
type InMemoryEmployeeRepository struct {
	employees map[int]*employee.Employee
	mu        sync.RWMutex
}

func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		employees: make(map[int]*employee.Employee),
	}
}

func (repo *InMemoryEmployeeRepository) CreateEmployee(emp *employee.Employee) *customerrors.RestErr {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.employees[emp.ID]; exists {
		errResponse := customerrors.RestErr{
			Message:   alreadyExists,
			ErrorCode: "ERR-R001",
		}
		return &errResponse
	}
	repo.employees[emp.ID] = emp
	return nil
}

func (repo *InMemoryEmployeeRepository) GetEmployeeByID(id int) (*employee.Employee, *customerrors.RestErr) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	emp, exists := repo.employees[id]
	if !exists {
		errResponse := customerrors.RestErr{
			Message:   notFound,
			ErrorCode: "ERR-R002",
		}
		return nil, &errResponse
	}
	return emp, nil
}

func (repo *InMemoryEmployeeRepository) UpdateEmployee(emp *employee.Employee) *customerrors.RestErr {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.employees[emp.ID]; !exists {
		errResponse := customerrors.RestErr{
			Message:   notFound,
			ErrorCode: "ERR-R003",
		}
		return &errResponse
	}
	repo.employees[emp.ID] = emp
	return nil
}

func (repo *InMemoryEmployeeRepository) DeleteEmployee(id int) *customerrors.RestErr {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.employees[id]; !exists {
		errResponse := customerrors.RestErr{
			Message:   notFound,
			ErrorCode: "ERR-R004",
		}
		return &errResponse
	}
	delete(repo.employees, id)
	return nil
}

func (repo *InMemoryEmployeeRepository) ListEmployees(pageNumber, pageSize int) ([]*employee.Employee, *customerrors.RestErr) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	start := (pageNumber - 1) * pageSize
	end := start + pageSize

	var employees []*employee.Employee
	for _, emp := range repo.employees {
		employees = append(employees, emp)
	}

	if start >= len(employees) {
		errResponse := customerrors.RestErr{
			Message:   incorrectPagination,
			ErrorCode: "ERR-R005",
		}
		return nil, &errResponse
	}

	if end > len(employees) {
		end = len(employees)
	}

	return employees[start:end], nil
}
