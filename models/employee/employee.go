package employee

type Employee struct {
	ID       int     `json:"id"  validate:"gt=0"`
	Name     string  `json:"name" validate:"required"`
	Position string  `json:"position" validate:"required"`
	Salary   float64 `json:"salary" validate:"gt=0"`
}
