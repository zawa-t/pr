package bitbucket

import "time"

type PullRequestComments struct {
	Size     int       `json:"size"`
	Page     int       `json:"page"`
	Pagelen  int       `json:"pagelen"`
	Next     string    `json:"next,omitempty"`
	Previous string    `json:"previous"`
	Values   []Comment `json:"values"`
}

type Comment struct {
	ID        int       `json:"id"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	Content   Content   `json:"content"`
	User      struct {
		Type        string `json:"type"`
		DisplayName string `json:"display_name"`
		// Links Link
		UUID      string `json:"uuid"`
		AccountID string `json:"account_id"`
		Nickname  string `json:"nickname"`
	} `json:"user"`
	Deleted     bool   `json:"deleted"`
	Inline      Inline `json:"inline"`
	Type        string `json:"type"`
	PullRequest struct {
		Type  string `json:"type"`
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"pullrequest"`
	// Resolution  struct {
	// 	Type string `json:""`
	// }
	Pending bool `json:"pending"`
}
