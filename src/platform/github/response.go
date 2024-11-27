package github

import "time"

// GetPRCommentResponse ...
type GetPRCommentResponse struct {
	URL                 string `json:"url"`
	PullRequestReviewID int64  `json:"pull_request_review_id"`
	ID                  int    `json:"id"`
	NodeID              string `json:"node_id"`
	DiffHunk            string `json:"diff_hunk"`
	Path                string `json:"path"`
	CommitID            string `json:"commit_id"`
	OriginalCommitID    string `json:"original_commit_id"`
	User                struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		UserViewType      string `json:"user_view_type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Body              string    `json:"body"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	HTMLURL           string    `json:"html_url"`
	PullRequestURL    string    `json:"pull_request_url"`
	AuthorAssociation string    `json:"author_association"`
	Reactions         struct {
		URL        string `json:"url"`
		TotalCount int    `json:"total_count"`
		Num1       int    `json:"+1"`
		Num10      int    `json:"-1"`
		Laugh      int    `json:"laugh"`
		Hooray     int    `json:"hooray"`
		Confused   int    `json:"confused"`
		Heart      int    `json:"heart"`
		Rocket     int    `json:"rocket"`
		Eyes       int    `json:"eyes"`
	} `json:"reactions"`
	StartLine         int    `json:"start_line"`
	OriginalStartLine int    `json:"original_start_line"`
	StartSide         string `json:"start_side"`
	Line              int    `json:"line"`
	OriginalLine      int    `json:"original_line"`
	Side              string `json:"side"`
	OriginalPosition  int    `json:"original_position"`
	Position          int    `json:"position"`
	SubjectType       string `json:"subject_type"`
}
