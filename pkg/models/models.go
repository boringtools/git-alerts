package models

type GitHubUser struct {
	Username          string `json:"login"`
	Url               string `json:"url"`
	ProfileUrl        string `json:"html_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	Admin             bool   `json:"site_admin"`
}

type GitHubRepository struct {
	FullName    string `json:"full_name"`
	Private     bool   `json:"private"`
	HtmlURL     string `json:"html_url"`
	Description string `json:"description"`
	Fork        bool   `json:"fork"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Push        string `json:"pushed_at"`
	CloneURL    string `json:"clone_url"`
	Visiability string `json:"visibility"`
}

type GitHubPATLimits struct {
	Rate Limits `json:"rate"`
}

type Limits struct {
	Total     int `json:"limit"`
	Used      int `json:"used"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}

type GitHubAPIEndPoints struct {
	GetUsers string
}

type ReportFileNames struct {
	GitHubOrgUsers          string
	GitHubOrgPublicRepos    string
	GitHubOrgPublicReposNew string
}

type SlackPayload struct {
	Text string `json:"text"`
}
