package controls

import "fmt"

var employeeType = []string{
	"Part Time Employee",
	"Full Time Employee",
	"Back",
}

var partTimeEmployeeUpdating = []string{
	"Update ID",
	"Update First Name",
	"Update Last Name",
	"Update Birth Year",
	"Update Work Hours",
	"Update Salary Per Hour",
	"Back",
}

var fullTimeEmployeeUpdating = []string{
	"Update ID",
	"Update First Name",
	"Update Last Name",
	"Update Birth Year",
	"Update Salary",
	"Update Reward",
	"Back",
}

var menu = map[string][]string{
	"mainMenu": {
		"Add Employee",
		"List Employees",
		"Show Employee",
		"Update Employee",
		"Delete Employee",
		"Export File",
		"Import File",
		"Exit",
	},
	"employeeType": employeeType,
	"partTimeEmployeeUpdating": partTimeEmployeeUpdating,
	"fullTimeEmployeeUpdating": fullTimeEmployeeUpdating,
}

func DisplayMenu(menuType string) {
	for index, value := range menu[menuType] {
		fmt.Println(fmt.Sprint(index + 1) + ".", value)
	}
}
