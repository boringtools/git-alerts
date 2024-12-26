package common

import (
	"encoding/csv"
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

func GetCSVFileContent(filePath string) ([]string, error) {
	openFile, errOpenFile := os.Open(filePath)

	if errOpenFile != nil {
		return nil, errOpenFile
	}

	defer openFile.Close()

	reader := csv.NewReader(openFile)
	records, errRecords := reader.ReadAll()

	if errRecords != nil {
		return nil, errRecords
	}

	var lines []string

	for _, record := range records {
		lines = append(lines, record[0])
	}

	return lines, nil
}
