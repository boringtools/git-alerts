package common

import (
	"os"
)

func SaveToJson(content []byte, filePath string) error {

	errWriteFile := os.WriteFile(filePath, content, 0644)

	if errWriteFile != nil {
		return errWriteFile
	}

	return nil
}
