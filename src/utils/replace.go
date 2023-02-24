package utils

import (
	"os"
	"strings"
)

// Replace -
func Replace(templatePath, tag, value string) (res string, err error) {
	file, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	return strings.Replace(string(file), tag, value, -1), nil
}
