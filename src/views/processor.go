package views

import (
	"employee_management/src/controls"
	"employee_management/src/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func initialHuman(array AnyArray, human HumanPointer) {
	var id string
	var birthYear int
	for {
		id = controls.InputStrictly("Enter ID: ")
		if !controls.CheckIDExists(array, id) {
			break
		} else {
			controls.PrintRedText("ID Exists")
		}
	}
	firstName := controls.InputStrictly("Enter First Name: ")
	lastName := controls.InputStrictly("Enter Last Name: ")
	for {
		birthYear = controls.ConvertStringToInteger(controls.InputStrictly("Enter Birth Year: "))
		if controls.CheckBirthYear(birthYear) {
			break
		} else {
			controls.PrintRedText("Invalid Birth Year")
		}
	}
	*human = models.Human{
		Id: id,
		FirstName: firstName,
		LastName: lastName,
		BirthYear: birthYear,
	}
}

func addPartTimeEmployee(human HumanPointer, partTimeEmployee PartTimeEmployeePointer) {
	var workHours float64
	var salaryPerHour float64
	for {
		workHours = controls.ConvertStringToFloat(controls.InputStrictly("Enter Work Hours: "))
		if controls.CheckNumbers(workHours) {
			break
		} else {
			controls.PrintRedText("Invalid Work Hours")
		}
	}
	for {
		salaryPerHour = controls.ConvertStringToFloat(controls.InputStrictly("Enter Salary Per Hour: "))
		if controls.CheckNumbers(salaryPerHour) {
			break
		} else {
			controls.PrintRedText("Invalid Salary Per Hour")
		}
	}
	*partTimeEmployee = models.PartTimeEmployee{
		Human: *human,
		WorkHours: workHours,
		SalaryPerHour: salaryPerHour,
	}
}

func addFullTimeEmployee(human HumanPointer, fullTimeEmployee FullTimeEmployeePointer) {
	var salary float64
	var reward float64
	for {
		salary = controls.ConvertStringToFloat(controls.Input("Enter Salary: "))
		if controls.CheckNumbers(salary) {
			break
		} else {
			controls.PrintRedText("Invalid Salary")
		}
	}
	for {
		reward = controls.ConvertStringToFloat(controls.Input("Enter Reward: "))
		if controls.CheckNumbers(reward) {
			break
		} else {
			controls.PrintRedText("Invalid Reward")
		}
	}
	*fullTimeEmployee = models.FullTimeEmployee{
		Human: *human,
		Salary: salary,
		Reward: reward,
	}
}

func IsEmpty(array AnyArray) bool {
	return len(array) == 0
}

func SearchEntity(array AnyArray, id string) (int, PartTimeEmployeePointer, FullTimeEmployeePointer) {
	var isFound bool
	for index, value := range array {
		switch employeeType := value.(type) {
		case *models.PartTimeEmployee:
			if employeeType.Id == id {
				isFound = true
				return index, employeeType, nil
			}
		case *models.FullTimeEmployee:
			if employeeType.Id == id {
				isFound = true
				return index, nil, employeeType
			}
		}
	}
	if !isFound {
		controls.PrintRedText("Employee Not Found")
	}
	return -1, nil, nil
}

func UpdateEntity(array AnyArray, employee interface{}, field string) {
	targetField := strings.ToLower(field)
	for {
		newValue := controls.InputStrictly("Enter New " + field + ": ")
		switch employeeType := employee.(type) {
		case PartTimeEmployeePointer:
			if targetField == "id" {
				if !controls.CheckIDExists(employeeArray, newValue) {
					employeeType.Id = newValue
				} else {
					controls.PrintRedText("ID Exists")
					continue
				}
			} else if targetField == "first name" {
				employeeType.FirstName = newValue
			} else if targetField == "last name" {
				employeeType.LastName = newValue
			} else if targetField == "birth year" {
				birthYear := controls.ConvertStringToInteger(newValue)
				if controls.CheckBirthYear(birthYear) {
					employeeType.BirthYear = birthYear
				} else {
					controls.PrintRedText("Invalid Birth Year")
					continue
				}
			} else if targetField == "work hours" {
				workHours := controls.ConvertStringToFloat(newValue)
				if controls.CheckNumbers(workHours) {
					employeeType.WorkHours = workHours
				} else {
					controls.PrintRedText("Invalid Work Hours")
					continue
				}
			} else if targetField == "salary per hour" {
				salaryPerHour := controls.ConvertStringToFloat(newValue)
				if controls.CheckNumbers(salaryPerHour) {
					employeeType.SalaryPerHour = salaryPerHour
				} else {
					controls.PrintRedText("Invalid Salary Per Hour")
					continue
				}
			}
		case FullTimeEmployeePointer:
			if targetField == "id" {
				if !controls.CheckIDExists(employeeArray, newValue) {
					employeeType.Id = newValue
				} else {
					controls.PrintRedText("ID Exists")
					continue
				}
			} else if targetField == "first name" {
				employeeType.FirstName = newValue
			} else if targetField == "last name" {
				employeeType.LastName = newValue
			} else if targetField == "birth year" {
				birthYear := controls.ConvertStringToInteger(newValue)
				if controls.CheckBirthYear(birthYear) {
					employeeType.BirthYear = birthYear
				} else {
					fmt.Println("Invalid Birth Year")
					continue
				}
			} else if targetField == "salary" {
				salary := controls.ConvertStringToFloat(newValue)
				if controls.CheckNumbers(salary) {
					employeeType.Salary = salary
				} else {
					controls.PrintRedText("Invalid Salary")
					continue
				}
			} else if targetField == "reward" {
				reward := controls.ConvertStringToFloat(newValue)
				if controls.CheckNumbers(reward) {
					employeeType.Reward = reward
				} else {
					controls.PrintRedText("Invalid Reward")
					continue
				}
			}
		}
		break
	}
	controls.PrintGreenText("Updated")
}

func DeleteEntity(array AnyArray, index int) AnyArray {
	return append(array[:index], array[index + 1:]...)
}

func ExportEntity(fileName string, array AnyArray) {
	if !strings.HasSuffix(fileName, ".json") {
		controls.PrintRedText(fileName + " must has .json extension")
		return
	}
	fileName = strings.ReplaceAll(fileName, " ", "_")
	json, err := json.MarshalIndent(array, "", "    ")
	if err != nil {
		controls.PrintRedText(fmt.Sprint(err))
		return
	}
	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		controls.PrintRedText(fmt.Sprint(err))
		return
	}
	controls.PrintGreenText("Exported")
}

func ImportEntity(path string, array *[]interface{}) []interface{} {
	data, err := os.ReadFile(path)
	if err != nil {
		controls.PrintRedText("Fail to access file with path " + path)
		return []interface{}{}
	}
	var temporaryArray []map[string]interface{}
	err = json.Unmarshal(data, &temporaryArray)
	if err != nil {
		controls.PrintRedText("Fail to convert to JSON")
		return []interface{}{}
	}
	var returnedArray []interface{}
	for _, entity := range temporaryArray {
		idValue, idOk := entity["id"]
		firstNameValue, firstNameOk := entity["first_name"]
		lastNameValue, lastNameOk := entity["last_name"]
		birthYearValue, birthYearOk := entity["birth_year"]
		workHoursValue, workHoursOk := entity["work_hours"]
		salaryPerHourValue, salaryPerHourOk := entity["salary_per_hour"]
		salaryValue, salaryOk := entity["salary"]
		rewardValue, rewardOk := entity["reward"]
		var birthYear int
		if !idOk {
			controls.PrintYellowText(fmt.Sprint(entity) + " don't have id, skipped")
			continue
		} else {
			if !controls.IsDesiredDataType(idValue, "string") {
				controls.PrintYellowText(fmt.Sprint(entity) + " wrong id data type, expected string, skipped")
				continue
			} else {
				if controls.CheckIDExists(*array, idValue.(string)) {
					controls.PrintYellowText(fmt.Sprint(entity) + " id already exists, skipped")
					continue
				}
			}
		}
		if !firstNameOk {
			controls.PrintYellowText(fmt.Sprint(entity) + " don't have first_name, skipped")
			continue
		} else {
			if !controls.IsDesiredDataType(firstNameValue, "string") {
				controls.PrintYellowText(fmt.Sprint(entity) + " wrong first_name data type, expected string, skipped")
				continue
			}
		}
		if !lastNameOk {
			controls.PrintYellowText(fmt.Sprint(entity) + " don't have last_name, skipped")
			continue
		} else {
			if !controls.IsDesiredDataType(lastNameValue, "string") {
				controls.PrintYellowText(fmt.Sprint(entity) + " wrong last_name data type, expected string, skipped")
				continue
			}
		}
		if !birthYearOk {
			controls.PrintYellowText(fmt.Sprint(entity) + " don't have birth_year, skipped")
			continue
		} else {
			if !controls.IsDesiredDataType(birthYearValue, "float") {
				controls.PrintYellowText(fmt.Sprint(entity) + " wrong birth_year data type, expected integer, skipped")
				continue
			} else {
				birthYear = int(birthYearValue.(float64))
				if !controls.CheckBirthYear(birthYear) {
					controls.PrintYellowText(fmt.Sprint(entity) + " invalid birth_year, skipped")
					continue
				}
			}
		}
		if workHoursOk && salaryPerHourOk {
			if controls.IsDesiredDataType(workHoursValue, "float") && controls.IsDesiredDataType(salaryPerHourValue, "float") {
				if controls.CheckNumbers(workHoursValue.(float64)) && controls.CheckNumbers(salaryPerHourValue.(float64)) {
					returnedArray = append(returnedArray, &models.PartTimeEmployee{
						Human: models.Human{
							Id: idValue.(string),
							FirstName: firstNameValue.(string),
							LastName: lastNameValue.(string),
							BirthYear: birthYear,
						},
						WorkHours: workHoursValue.(float64),
						SalaryPerHour: salaryPerHourValue.(float64),
					})
					controls.PrintGreenText(fmt.Sprint(entity) + " imported")
				} else {
					controls.PrintYellowText(fmt.Sprint(entity) + " both work_hours & salary_per_hour must greater than 0")
					continue
				}
			} else {
				controls.PrintYellowText(fmt.Sprint(entity) + " wrong work_hours & salary_per_hour data type, both expected float, skipped")
				continue
			}
		} else if salaryOk && rewardOk {
			if controls.IsDesiredDataType(salaryValue, "float") && controls.IsDesiredDataType(rewardValue, "float") {
				if controls.CheckNumbers(salaryValue.(float64)) && controls.CheckNumbers(rewardValue.(float64)) {
					returnedArray = append(returnedArray, &models.FullTimeEmployee{
						Human: models.Human{
							Id: idValue.(string),
							FirstName: firstNameValue.(string),
							LastName: lastNameValue.(string),
							BirthYear: birthYear,
						},
						Salary: salaryValue.(float64),
						Reward: rewardValue.(float64),
					})
					controls.PrintGreenText(fmt.Sprint(entity) + " imported")
				} else {
					controls.PrintYellowText(fmt.Sprint(entity) + " both salary & reward must greater than 0")
					continue
				}
			} else {
				controls.PrintYellowText(fmt.Sprint(entity) + " wrong salary & reward data type, both expected float, skipped")
				continue
			}
		} else {
			controls.PrintYellowText(fmt.Sprint(entity) + " invalid employee type, skipped")
			continue
		}
	}
	return returnedArray
}
