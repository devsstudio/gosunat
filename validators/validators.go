package validators

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/devsstudio/gosunat/enums"
	"github.com/devsstudio/gosunat/helpers"
)

func AlphanumericStartWith(value, startWith string, min, max int) bool {
	return strings.HasPrefix(value, startWith) &&
		govalidator.IsAlphanumeric(value) &&
		len(value) >= min && len(value) <= max
}

func IsLength(value string, min, max int) bool {
	return govalidator.StringLength(value, strconv.Itoa(min), strconv.Itoa(max))
}

func IsIntGreaterThan(value string, gt int) bool {
	intValue, error := strconv.Atoi(value)
	if error == nil {
		return intValue > gt
	} else {
		return false
	}
}

func IsSale(fileCode string) bool {
	return helpers.Contains([]string{
		string(enums.AF_FILE_14),
		string(enums.AF_FILE_14_1),
		string(enums.AF_FILE_14_2),
		string(enums.PLE_CODE_14_1),
		string(enums.PLE_CODE_14_2),
	}, fileCode)
}

func IsPurchase(fileCode string) bool {
	return helpers.Contains([]string{
		string(enums.AF_FILE_8),
		string(enums.AF_FILE_8_1),
		string(enums.AF_FILE_8_2),
		string(enums.AF_FILE_8_3),
		string(enums.PLE_CODE_8_1),
		string(enums.PLE_CODE_8_2),
		string(enums.PLE_CODE_8_3),
	}, fileCode)
}

// Determina si un código de libro es venta o compra
func IsSaleOrPurchase(fileCode string) bool {
	return IsSale(fileCode) || IsPurchase(fileCode)
}

func IsNotAllZeroes(s string) bool {
	matched, _ := regexp.MatchString(`^0+$`, s)
	return !matched
}

//TODO: SI govalidator.IsInt no funciona, usar esta
/*func IsInt(s string, allowLeadingZeroes bool) bool {
	if !allowLeadingZeroes {
		matched, _ := regexp.MatchString(`^[1-9]\d*$`, s)
		return matched
	}
	_, err := strconv.Atoi(s)
	return err == nil
}*/
