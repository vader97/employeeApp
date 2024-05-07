package api

import (
	"employeeApp/models/customerrors"
	"employeeApp/models/employee"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"employeeApp/service"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		service: service,
	}
}

func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rerr := customerrors.RestErr{
			Message:   "Invalid employee ID",
			ErrorCode: "ERR-H001",
		}
		c.JSON(http.StatusBadRequest, rerr)
		return
	}

	emp, rerr := h.service.GetEmployeeByID(id)
	if rerr != nil {
		c.JSON(http.StatusNotFound, rerr)
		return
	}

	c.JSON(http.StatusOK, emp)
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var emp employee.Employee
	if err := c.BindJSON(&emp); err != nil {
		rerr := customerrors.RestErr{
			Message:   "Invalid employee data",
			ErrorCode: "ERR-H002",
		}
		c.JSON(http.StatusBadRequest, &rerr)
		return
	}
	validate := validator.New()
	if err := validate.Struct(emp); err != nil {
		rerr := customerrors.RestErr{
			Message:   "Invalid employee data",
			ErrorCode: "ERR-H003",
		}
		c.JSON(http.StatusBadRequest, rerr)
		return
	}

	if err := h.service.CreateEmployee(&emp); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rerr := customerrors.RestErr{
			Message:   "Invalid employee ID",
			ErrorCode: "ERR-H004",
		}
		c.JSON(http.StatusBadRequest, rerr)
		return
	}

	var emp employee.Employee
	if err := c.BindJSON(&emp); err != nil {
		rerr := customerrors.RestErr{
			Message:   "Invalid employee data",
			ErrorCode: "ERR-H005",
		}
		c.JSON(http.StatusBadRequest, rerr)
		return
	}
	validate := validator.New()
	if err := validate.Struct(emp); err != nil {
		rerr := customerrors.RestErr{
			Message:   "Invalid employee data",
			ErrorCode: "ERR-H006",
		}
		c.JSON(http.StatusBadRequest, rerr)
		return
	}
	emp.ID = id

	if rerr := h.service.UpdateEmployee(&emp); rerr != nil {
		c.JSON(http.StatusInternalServerError, rerr)
		return
	}

	c.Status(http.StatusOK)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rerr := customerrors.RestErr{
			Message:   "Invalid employee ID",
			ErrorCode: "ERR-H007",
		}
		c.JSON(http.StatusBadRequest, rerr)
		return
	}

	if rerr := h.service.DeleteEmployee(id); rerr != nil {
		c.JSON(http.StatusInternalServerError, rerr)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *EmployeeHandler) ListEmployees(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		rerr := customerrors.RestErr{
			Message:   "Invalid page number",
			ErrorCode: "ERR-H008",
		}
		c.JSON(http.StatusBadRequest, rerr)
		return
	}

	// default page size is 10
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size <= 0 {
		rerr := customerrors.RestErr{
			Message:   "Invalid page size",
			ErrorCode: "ERR-H009",
		}
		c.JSON(http.StatusBadRequest, rerr)
		return
	}

	employees, rerr := h.service.ListEmployees(page, size)
	if rerr != nil {
		c.JSON(http.StatusInternalServerError, rerr)
		return
	}

	c.JSON(http.StatusOK, employees)
}
