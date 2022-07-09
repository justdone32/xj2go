package demojson2

import "time"

type Github struct {
	CloneURL         string       `json:"clone_url"`
	CreatedAt        time.Time    `json:"created_at"`
	DefaultBranch    string       `json:"default_branch"`
	Description      string       `json:"description"`
	Fork             bool         `json:"fork"`
	ForksCount       float64      `json:"forks_count"`
	FullName         string       `json:"full_name"`
	GitURL           string       `json:"git_url"`
	HTMLURL          string       `json:"html_url"`
	HasDownloads     bool         `json:"has_downloads"`
	HasIssues        bool         `json:"has_issues"`
	HasWiki          bool         `json:"has_wiki"`
	Homepage         string       `json:"homepage"`
	ID               float64      `json:"id"`
	MirrorURL        string       `json:"mirror_url"`
	Name             string       `json:"name"`
	OpenIssuesCount  float64      `json:"open_issues_count"`
	Organization     Organization `json:"organization"`
	Owner            Owner        `json:"owner"`
	Parent           Parent       `json:"parent"`
	Permissions      Permissions  `json:"permissions"`
	Private          bool         `json:"private"`
	PushedAt         time.Time    `json:"pushed_at"`
	SSHURL           string       `json:"ssh_url"`
	Size             float64      `json:"size"`
	Source           Source       `json:"source"`
	StargazersCount  float64      `json:"stargazers_count"`
	SubscribersCount float64      `json:"subscribers_count"`
	SvnURL           string       `json:"svn_url"`
	URL              string       `json:"url"`
	UpdatedAt        time.Time    `json:"updated_at"`
	WatchersCount    float64      `json:"watchers_count"`
}
type Organization struct {
	AvatarURL         string  `json:"avatar_url"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        string  `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                float64 `json:"id"`
	Login             string  `json:"login"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
}
type Owner struct {
	AvatarURL         string  `json:"avatar_url"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        string  `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                float64 `json:"id"`
	Login             string  `json:"login"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
}
type Parent struct {
	CloneURL        string      `json:"clone_url"`
	CreatedAt       time.Time   `json:"created_at"`
	DefaultBranch   string      `json:"default_branch"`
	Description     string      `json:"description"`
	Fork            bool        `json:"fork"`
	ForksCount      float64     `json:"forks_count"`
	FullName        string      `json:"full_name"`
	GitURL          string      `json:"git_url"`
	HTMLURL         string      `json:"html_url"`
	HasDownloads    bool        `json:"has_downloads"`
	HasIssues       bool        `json:"has_issues"`
	HasWiki         bool        `json:"has_wiki"`
	Homepage        string      `json:"homepage"`
	ID              float64     `json:"id"`
	MirrorURL       string      `json:"mirror_url"`
	Name            string      `json:"name"`
	OpenIssuesCount float64     `json:"open_issues_count"`
	Owner           Owner       `json:"owner"`
	Permissions     Permissions `json:"permissions"`
	Private         bool        `json:"private"`
	PushedAt        time.Time   `json:"pushed_at"`
	SSHURL          string      `json:"ssh_url"`
	Size            float64     `json:"size"`
	StargazersCount float64     `json:"stargazers_count"`
	SvnURL          string      `json:"svn_url"`
	URL             string      `json:"url"`
	UpdatedAt       time.Time   `json:"updated_at"`
	WatchersCount   float64     `json:"watchers_count"`
}
type Permissions struct {
	Admin bool `json:"admin"`
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
}
type Source struct {
	CloneURL        string      `json:"clone_url"`
	CreatedAt       time.Time   `json:"created_at"`
	DefaultBranch   string      `json:"default_branch"`
	Description     string      `json:"description"`
	Fork            bool        `json:"fork"`
	ForksCount      float64     `json:"forks_count"`
	FullName        string      `json:"full_name"`
	GitURL          string      `json:"git_url"`
	HTMLURL         string      `json:"html_url"`
	HasDownloads    bool        `json:"has_downloads"`
	HasIssues       bool        `json:"has_issues"`
	HasWiki         bool        `json:"has_wiki"`
	Homepage        string      `json:"homepage"`
	ID              float64     `json:"id"`
	MirrorURL       string      `json:"mirror_url"`
	Name            string      `json:"name"`
	OpenIssuesCount float64     `json:"open_issues_count"`
	Owner           Owner       `json:"owner"`
	Permissions     Permissions `json:"permissions"`
	Private         bool        `json:"private"`
	PushedAt        time.Time   `json:"pushed_at"`
	SSHURL          string      `json:"ssh_url"`
	Size            float64     `json:"size"`
	StargazersCount float64     `json:"stargazers_count"`
	SvnURL          string      `json:"svn_url"`
	URL             string      `json:"url"`
	UpdatedAt       time.Time   `json:"updated_at"`
	WatchersCount   float64     `json:"watchers_count"`
}
