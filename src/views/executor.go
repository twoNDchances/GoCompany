package views

import (
	"employee_management/src/controls"
	"employee_management/src/models"
	"fmt"
	"os"
)

var employeeArray []interface{}

const fileNameExporting = "employees.json"

func Run() {
	fmt.Println("===== Welcome to Go Company =====")
	for {
		controls.DisplayMenu("mainMenu")
		selection := controls.ConvertStringToInteger(controls.Input("Your Selection: "))
		switch selection {
		case 1:
			var employeeType int
			for {
				controls.DisplayMenu("employeeType")
				employeeType = controls.ConvertStringToInteger(controls.Input("Select employee type: "))
				if employeeType <= 0 || employeeType > 3 {
					controls.PrintRedText("Invalid Employee Type, must in [1-3]")
				} else {
					break
				}
			}
			switch employeeType {
			case 1:
				var human models.Human
				var partTimeEmployee models.PartTimeEmployee
				initialHuman(employeeArray, &human)
				addPartTimeEmployee(&human, &partTimeEmployee)
				employeeArray = append(employeeArray, &partTimeEmployee)
			case 2:
				var human models.Human
				var fullTimeEmployee models.FullTimeEmployee
				initialHuman(employeeArray, &human)
				addFullTimeEmployee(&human, &fullTimeEmployee)
				employeeArray = append(employeeArray, &fullTimeEmployee)
			case 3:
				continue
			}
		case 2:
			if IsEmpty(employeeArray) {
				controls.PrintRedText("No Employee Data")
				continue
			}
			for _, employee := range employeeArray {
				models.PrintEntity(employee.(models.StringOuter))
			}
		case 3:
			if IsEmpty(employeeArray) {
				controls.PrintRedText("No Employee Data")
				continue
			}
			id := controls.InputStrictly("Enter ID: ")
			_, partTime, fullTime := SearchEntity(employeeArray, id)
			if partTime != nil {
				models.PrintEntity(*partTime)
			}
			if fullTime != nil {
				models.PrintEntity(*fullTime)
			}
		case 4:
			if IsEmpty(employeeArray) {
				controls.PrintRedText("No Employee Data")
				continue
			}
			id := controls.InputStrictly("Enter ID: ")
			_, partTime, fullTime := SearchEntity(employeeArray, id)
			var updatingSelection int
			if partTime != nil {
				for {
					controls.DisplayMenu("partTimeEmployeeUpdating")
					updatingSelection = controls.ConvertStringToInteger(controls.Input("Your updating selection: "))
					if updatingSelection <= 0 || updatingSelection > 7 {
						controls.PrintRedText("Updating Selection must in [1-7]")
					} else {
						break
					}
				}
				switch updatingSelection {
				case 1:
					UpdateEntity(employeeArray, partTime, "ID")
				case 2:
					UpdateEntity(employeeArray, partTime, "First Name")
				case 3:
					UpdateEntity(employeeArray, partTime, "Last Name")
				case 4:
					UpdateEntity(employeeArray, partTime, "Birth Year")
				case 5:
					UpdateEntity(employeeArray, partTime, "Work Hours")
				case 6:
					UpdateEntity(employeeArray, partTime, "Salary Per Hour")
				case 7:
					continue
				}
			}
			if fullTime != nil {
				for {
					controls.DisplayMenu("fullTimeEmployeeUpdating")
					updatingSelection = controls.ConvertStringToInteger(controls.Input("Your updating selection: "))
					if updatingSelection <= 0 || updatingSelection > 7 {
						controls.PrintRedText("Updating Selection must in [1-7]")
					} else {
						break
					}
				}
				switch updatingSelection {
				case 1:
					UpdateEntity(employeeArray, fullTime, "ID")
				case 2:
					UpdateEntity(employeeArray, fullTime, "First Name")
				case 3:
					UpdateEntity(employeeArray, fullTime, "Last Name")
				case 4:
					UpdateEntity(employeeArray, fullTime, "Birth Year")
				case 5:
					UpdateEntity(employeeArray, fullTime, "Salary")
				case 6:
					UpdateEntity(employeeArray, fullTime, "Reward")
				case 7:
					break
				}
			}
		case 5:
			if IsEmpty(employeeArray) {
				controls.PrintRedText("No Employee Data")
				continue
			}
			id := controls.InputStrictly("Enter ID: ")
			index, _, _ := SearchEntity(employeeArray, id)
			if index == -1 {
				continue
			}
			employeeArray = DeleteEntity(employeeArray, index)
			controls.PrintGreenText("Deleted")
		case 6:
			if IsEmpty(employeeArray) {
				controls.PrintRedText("No Employee Data")
				continue
			}
			ExportEntity(fileNameExporting, employeeArray)
		case 7:
			path := controls.InputStrictly("Enter Path: ")
			employeeArray = append(employeeArray, ImportEntity(path, &employeeArray)...)
		case 8:
			controls.PrintGreenText("Good bye!!!")
			os.Exit(0)
		default:
			controls.PrintRedText("Invalid Selection, must in [1-8]")
		}
	}
}
