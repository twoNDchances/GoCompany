package models

import "fmt"

type PartTimeEmployee struct {
	Human
	WorkHours float64 `json:"work_hours"`
	SalaryPerHour float64 `json:"salary_per_hour"`
}

func (partTimeEmployee PartTimeEmployee) ToString() string {
	return partTimeEmployee.Human.ToString() + ", WorkHours: " + fmt.Sprint(partTimeEmployee.WorkHours) + ", SalaryPerHour: " + fmt.Sprint(partTimeEmployee.SalaryPerHour)
}
