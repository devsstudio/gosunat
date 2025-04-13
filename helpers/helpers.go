package helpers

import "fmt"

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsAsString[T comparable](slice []T, value string) bool {
	for _, v := range slice {
		if fmt.Sprintf("%v", v) == value {
			return true
		}
	}
	return false
}
