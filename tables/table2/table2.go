package table2

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Table2 string

const (
	ID_1 Table2 = "1"
	ID_4 Table2 = "4"
	ID_6 Table2 = "6"
	ID_7 Table2 = "7"
	ID_A Table2 = "A"
	ID_0 Table2 = "0"
)

var TABLE_2_ITEMS = map[Table2]string{
	ID_1: "DNI",
	ID_4: "Carnet de extranjería",
	ID_6: "RUC",
	ID_7: "Pasaporte",
	ID_A: "Cedula diplomática de identidad",
	ID_0: "Otro",
}

func ValidateCode(code string) bool {
	allowedCodes := []Table2{ID_1, ID_4, ID_6, ID_7, ID_A, ID_0}
	for _, validCode := range allowedCodes {
		if Table2(code) == validCode {
			return true
		}
	}
	return false
}

func ValidateNumber(code Table2, number string) bool {
	validate := validator.New()

	switch code {
	case ID_1:
		// DNI: exactamente 8 dígitos
		isInt, _ := regexp.MatchString(`^\d{8}$`, number)
		return isInt

	case ID_4, ID_7, ID_A:
		// Alfanumérico entre 1 y 12 caracteres
		return validate.Var(number, "alphanum,min=1,max=12") == nil

	case ID_6:
		// RUC: exactamente 11 dígitos
		isInt, _ := regexp.MatchString(`^\d{11}$`, number)
		return isInt

	case ID_0:
		// Alfanumérico entre 1 y 15 caracteres
		return validate.Var(number, "alphanum,min=1,max=15") == nil

	default:
		return false
	}
}

func GetName(value string) string {
	return TABLE_2_ITEMS[Table2(value)]
}
