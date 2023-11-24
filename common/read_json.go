package common

import (
	"io"
	"os"

	"github.com/boringtools/git-alerts/logger"
)

func GetJsonFileContent(fileName string) (jsonData []byte) {
	filePath := os.Getenv("rfp") + fileName + ".json"

	file, errFile := os.Open(filePath)

	if errFile != nil {
		logger.LogERR("GetJsonFileContent - Error in location file")
	}

	jsonData, errRead := io.ReadAll(file)

	if errRead != nil {
		logger.LogERR("GetJsonFileContent - Error in reading file content")
	}

	return jsonData
}
