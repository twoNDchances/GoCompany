package models

import "fmt"

type Human struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthYear int `json:"birth_year"`
}

func (human Human) ToString() string {
	return "Id: " + human.Id + ", FirstName: " + human.FirstName + ", LastName: " + human.LastName + ", BirthYear: " + fmt.Sprint(human.BirthYear)
}
