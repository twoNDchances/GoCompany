package models

import (
	"fmt"
)

type FullTimeEmployee struct {
	Human
	Salary float64 `json:"salary"`
	Reward float64 `json:"reward"`
}

func (fullTimeEmployee FullTimeEmployee) ToString() string {
	return fullTimeEmployee.Human.ToString() + ", Salary: " + fmt.Sprint(fullTimeEmployee.Salary) + ", Reward: " + fmt.Sprint(fullTimeEmployee.Reward)
}
