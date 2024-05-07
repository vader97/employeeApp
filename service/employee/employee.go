package employee

import (
	"employeeApp/models/customerrors"
	"employeeApp/models/employee"
	"employeeApp/repository"
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
			Message:   "invalid employee data",
			ErrorCode: "ERR-S001",
		}
		return &errResponse
	}
	return s.repo.CreateEmployee(emp)
}

func (s *EmployeeServiceImpl) GetEmployeeByID(id int) (*employee.Employee, *customerrors.RestErr) {
	if id <= 0 {
		errResponse := customerrors.RestErr{
			Message:   "invalid employee ID",
			ErrorCode: "ERR-S002",
		}
		return nil, &errResponse
	}
	return s.repo.GetEmployeeByID(id)
}

func (s *EmployeeServiceImpl) UpdateEmployee(emp *employee.Employee) *customerrors.RestErr {
	if emp == nil || emp.ID <= 0 {
		errResponse := customerrors.RestErr{
			Message:   "invalid employee data",
			ErrorCode: "ERR-S003",
		}
		return &errResponse
	}
	return s.repo.UpdateEmployee(emp)
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) *customerrors.RestErr {
	if id <= 0 {
		errResponse := customerrors.RestErr{
			Message:   "invalid employee ID",
			ErrorCode: "ERR-S004",
		}
		return &errResponse
	}
	return s.repo.DeleteEmployee(id)
}

func (s *EmployeeServiceImpl) ListEmployees(pageNumber, pageSize int) ([]*employee.Employee, *customerrors.RestErr) {
	if pageNumber <= 0 || pageSize <= 0 {
		errResponse := customerrors.RestErr{
			Message:   "invalid page number or page size",
			ErrorCode: "ERR-S005",
		}
		return nil, &errResponse
	}
	return s.repo.ListEmployees(pageNumber, pageSize)
}
