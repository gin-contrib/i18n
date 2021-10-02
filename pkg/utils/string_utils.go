package utils

import (
	"strings"
)

const (
	UndefinedString   = "undefined"
	ZeroString = ""
	CensorStringValue = "***"
)

// IsStringSliceContains -- check slice contain string
func IsStringSliceContains(stringSlice []string, searchString string) bool {
	for _, value := range stringSlice {
		if value == searchString {
			return true
		}
	}
	return false
}

// StringTrimSpace -- trim space of string
func StringTrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// IsStringEmpty -- check if string is empty
func IsStringEmpty(s string) bool {
	return s == ZeroString
}

// IsStringNotEmpty -- check if string is not empty
func IsStringNotEmpty(s string) bool {
	return s != ZeroString
}

// CensorString --
func CensorString(str string) string {
	if len(str) <= 6 {
		return CensorStringValue
	}

	return str[:2] + CensorStringValue + str[len(str)-2:]
}

// StringPrefixInSlice --
func StringPrefixInSlice(str string, list []string) bool {
	for _, v := range list {
		if strings.HasPrefix(str, v) {
			return true
		}
	}
	return false
}

// StringInSliceEqualFold --
func StringInSliceEqualFold(str string, list []string) bool {
	for _, v := range list {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
}

// IndexInSliceString --
func IndexInSliceString(str string, list []string) int {
	for index, v := range list {
		if v == str {
			return index
		}
	}
	return -1
}

// DeleteElementInSliceString --
func DeleteElementInSliceString(list []string, index int) []string {
	list = append(list[:index], list[index+1:]...)
	return list
}

// StringKeys --
func StringKeys(mmap map[string]interface{}) []string {
	keys := make([]string, 0, len(mmap))
	for k := range mmap {
		keys = append(keys, k)
	}
	return keys
}

// IsUndefinedValue
func IsUndefinedValue(s string) bool {
	return strings.EqualFold(s, UndefinedString)
}