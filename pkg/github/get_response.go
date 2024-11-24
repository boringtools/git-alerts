package github

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var (
	pageLength int
)

func GetGitHubResponse(URL string, auth bool, params map[string]string) ([]byte, int, error) {

	parameters := url.Values{}

	for key, value := range params {
		parameters.Add(key, value)
	}

	fullURL := fmt.Sprintf("%s?%s", URL, parameters.Encode())
	client := &http.Client{}

	request, errRequest := http.NewRequest("GET", fullURL, nil)

	if errRequest != nil {
		return nil, 0, errRequest
	}

	if auth {
		token := "Bearer " + os.Getenv("GITHUB_PAT")
		request.Header.Add("Authorization", token)
	}

	response, errResponse := client.Do(request)

	if errResponse != nil {
		return nil, 0, errResponse
	}

	if response.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("unable to fetch response, status code : %s", response.Status)
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
		return nil, 0, errRead
	}
	defer response.Body.Close()

	return rowResponse, pageLength, nil
}
