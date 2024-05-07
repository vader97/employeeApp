package employee

import (
	"employeeApp/models/customerrors"
	"employeeApp/models/employee"
	"employeeApp/repository"
	"net/http"
)

type EmployeeServiceImpl struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		repo: repo,
	}
}

func (s *EmployeeServiceImpl) CreateEmployee(emp *employee.Employee) *customerrors.RestErr {
	if emp == nil {
		errResponse := customerrors.RestErr{
			Message:    invalidData,
			ErrorCode:  "ERR-S001",
			StatusCode: http.StatusBadRequest,
		}
		return &errResponse
	}
	return s.repo.CreateEmployee(emp)
}

func (s *EmployeeServiceImpl) GetEmployeeByID(id int) (*employee.Employee, *customerrors.RestErr) {
	if id <= 0 {
		errResponse := customerrors.RestErr{
			Message:    invalidId,
			ErrorCode:  "ERR-S002",
			StatusCode: http.StatusBadRequest,
		}
		return nil, &errResponse
	}
	return s.repo.GetEmployeeByID(id)
}

func (s *EmployeeServiceImpl) UpdateEmployee(emp *employee.Employee) *customerrors.RestErr {
	if emp == nil || emp.ID <= 0 {
		errResponse := customerrors.RestErr{
			Message:    invalidData,
			ErrorCode:  "ERR-S003",
			StatusCode: http.StatusBadRequest,
		}
		return &errResponse
	}
	return s.repo.UpdateEmployee(emp)
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) *customerrors.RestErr {
	if id <= 0 {
		errResponse := customerrors.RestErr{
			Message:    invalidId,
			ErrorCode:  "ERR-S004",
			StatusCode: http.StatusBadRequest,
		}
		return &errResponse
	}
	return s.repo.DeleteEmployee(id)
}

func (s *EmployeeServiceImpl) ListEmployees(pageNumber, pageSize int) ([]*employee.Employee, *customerrors.RestErr) {
	if pageNumber <= 0 || pageSize <= 0 {
		errResponse := customerrors.RestErr{
			Message:    invalidPagination,
			ErrorCode:  "ERR-S005",
			StatusCode: http.StatusBadRequest,
		}
		return nil, &errResponse
	}
	return s.repo.ListEmployees(pageNumber, pageSize)
}
