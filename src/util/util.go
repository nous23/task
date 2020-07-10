package util

import (
	"os"
	"strings"
)

// report whether path exist
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteLast(s, d string) string {
	index := strings.LastIndex(s, d)
	if index == -1 {
		return s
	}
	return s[:index] + s[index+1:]
}

func HasKey(m map[string]interface{}, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

func IsEmpty(m map[string]interface{}) bool {
	return len(m) == 0
}
