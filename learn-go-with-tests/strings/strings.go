package strings

import (
	"strings"
)

func CaseConverter(inputString string, requiredCase string) string {
	var returnString string
	switch requiredCase {
	case "lower":
		returnString = strings.ToLower(inputString)
	case "upper":
		returnString = strings.ToUpper(inputString)
	}

	return returnString
}
