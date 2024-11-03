package bitbucket

// NOTE: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type CommentData struct {
	Content Content `json:"content,omitempty"`
	Inline  Inline  `json:"inline,omitempty"`
}

type Content struct {
	Raw    string `json:"raw,omitempty"`
	Markup string `json:"markup,omitempty"`
	HTML   string `json:"html,omitempty"`
}

type Inline struct {
	Path string `json:"path,omitempty"`
	To   uint   `json:"to,omitempty"`
}
