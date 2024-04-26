package gh

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/boringtools/git-alerts/logger"
)

var (
	pageLength int
)

func GetResponse(ghURL string, auth bool, params map[string]string) ([]byte, int) {

	parameters := url.Values{}

	for key, value := range params {
		parameters.Add(key, value)
	}

	fullURL := fmt.Sprintf("%s?%s", ghURL, parameters.Encode())
	client := &http.Client{}

	request, errRequest := http.NewRequest("GET", fullURL, nil)

	if errRequest != nil {
		logger.LogERR("GetResponse - Error in making HTTP calls to GitHub")
		panic(errRequest)
	}

	if auth {
		pat := "token " + os.Getenv("GITHUB_PAT")
		request.Header.Add("Authorization", pat)
	}

	response, errResponse := client.Do(request)

	if errResponse != nil {
		logger.LogERR("GetResponse - Error in fetching HTTP Response")
		panic(errResponse)
	}

	if response.StatusCode != 200 {
		logger.LogERR("Please provide a valid GitHub PAT")
		os.Exit(1)
	}

	getLinkAttr := response.Header.Get("Link")

	if getLinkAttr != "" {
		firstCut := strings.Split(getLinkAttr, ">;")
		seondCut := strings.Split(firstCut[1], "=")
		_, errPageLength := strconv.Atoi(seondCut[3])

		if errPageLength != nil {
			panic(errPageLength)
		} else {
			pageLength, _ = strconv.Atoi(seondCut[3])
		}

	} else {
		pageLength = 0
	}

	rowResponse, errRead := io.ReadAll(response.Body)

	if errRead != nil {
		logger.LogERR("GetResponse - Error in reading response body")
		panic(errRead)
	}
	defer response.Body.Close()

	return rowResponse, pageLength
}
