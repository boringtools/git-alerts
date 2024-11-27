package common

import (
	"io"
	"os"
)

func GetJSONFileContent(filePath string) ([]byte, error) {
	openFile, errOpenFile := os.Open(filePath)

	if errOpenFile != nil {
		return nil, errOpenFile
	}

	defer openFile.Close()

	byteData, errRead := io.ReadAll(openFile)

	if errRead != nil {
		return nil, errRead
	}

	return byteData, nil
}
