package github

import (
	"encoding/json"

	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/boringtools/git-alerts/pkg/models"
)

func GetGitHubPATLimits() (*models.Limits, error) {
	url := common.GitHubAPIRateLimit

	parameters := map[string]string{}
	var limit models.GitHubPATLimits

	ghResponse, _, _ := GetGitHubResponse(url, common.AuthenticatedScan, parameters)
	err := json.Unmarshal(ghResponse, &limit)

	if err != nil {
		return nil, err
	}

	return &models.Limits{
		Total:     limit.Rate.Total,
		Used:      limit.Rate.Used,
		Remaining: limit.Rate.Remaining,
		Reset:     limit.Rate.Reset,
	}, nil
}
