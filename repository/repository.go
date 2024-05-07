// repository.go in repository package
package repository

import (
	"employeeApp/models/customerrors"
	"employeeApp/models/employee"
)

type EmployeeRepository interface {
	CreateEmployee(emp *employee.Employee) *customerrors.RestErr
	GetEmployeeByID(id int) (*employee.Employee, *customerrors.RestErr)
	UpdateEmployee(emp *employee.Employee) *customerrors.RestErr
	DeleteEmployee(id int) *customerrors.RestErr
	ListEmployees(pageNumber, pageSize int) ([]*employee.Employee, *customerrors.RestErr)
}
