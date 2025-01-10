package controls

import (
	"employee_management/src/models"
	"errors"
	"time"
)

func checkDataType(data interface{}) (string, error) {
	switch data.(type) {
	case string:
		return "string", nil
	case int:
		return "int", nil
	case float64:
		return "float", nil
	case float32:
		return "float", nil
	default:
		return "", errors.New("invalid data type")
	}
}

func IsDesiredDataType(data interface{}, desiredDataType string) (bool) {
	checkDataType, err := checkDataType(data)
	if err != nil {
		return false
	}
	return checkDataType == desiredDataType
}

func CheckNumbers(data interface{}) bool {
	if IsDesiredDataType(data, "int") {
		return data.(int) > 0
	}
	if IsDesiredDataType(data, "float") {
		return data.(float64) > 0
	}
	return false
}

func CheckStrings(data interface{}) bool {
	if IsDesiredDataType(data, "string") {
		return len(data.(string)) > 0
	}
	return false
}

func CheckIDExists(data []interface{}, id string) bool {
	for _, employee := range data {
		switch employeeType := employee.(type) {
		case *models.PartTimeEmployee:
			if employeeType.Id == id {
				return true
			}
		case *models.FullTimeEmployee:
			if employeeType.Id == id {
				return true
			}
		}
	}
	return false
}

func CheckBirthYear(birthYear int) bool {
	return birthYear > 1900 && birthYear < time.Now().Year()
}
