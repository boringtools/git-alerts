package common

import (
	"os"

	"github.com/boringtools/git-alerts/logger"
)

func SaveToJson(content []byte, fileName string) {

	fullPath := os.Getenv("rfp") + fileName + ".json"

	os.WriteFile(fullPath, content, 0644)
	logger.LogP("Data saved successfully : ", fullPath)
}
