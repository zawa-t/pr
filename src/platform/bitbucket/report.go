package bitbucket

// NOTE: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-reports/#api-repositories-workspace-repo-slug-commit-commit-reports-reportid-put
type ReportData struct {
	UUID              string `json:"uuid,omitempty"`
	Title             string `json:"title,omitempty"`
	Details           string `json:"details,omitempty"`
	ExternalID        string `json:"external_id,omitempty"`
	Reporter          string `json:"reporter,omitempty"`
	Link              string `json:"link,omitempty"`
	RemoteLinkEnabled bool   `json:"remote_link_enabled,omitempty"`
	LogoURL           string `json:"logo_url,omitempty"`
	ReportType        string `json:"report_type,omitempty"` // SECURITY, COVERAGE, TEST, BUG
	Result            string `json:"result,omitempty"`      // PASSED, FAILED, PENDING
	Data              []Data `json:"data,omitempty"`
	CreatedOn         string `json:"created_on,omitempty"`
	UpdatedOn         string `json:"updated_on,omitempty"`
}

type Data struct {
	Type  string `json:"type"` // BOOLEAN, DATE, DURATION, LINK, NUMBER, PERCENTAGE, TEXT
	Title string `json:"title"`
	Value Value  `json:"value"`
}

type Value struct { // 好きなオブジェクトでOK
}

type AnnotationResponse struct {
	UUID string `json:"uuid,omitempty"`
}

type AnnotationData struct {
	UUID           string `json:"uuid,omitempty"`
	ExternalID     string `json:"external_id,omitempty"`
	Type           string `json:"type,omitempty"`
	AnnotationType string `json:"annotation_type,omitempty"` // VULNERABILITY, CODE_SMELL, BUG
	Path           string `json:"path,omitempty"`
	Line           uint   `json:"line,omitempty"`
	Summary        string `json:"summary,omitempty"`
	Details        string `json:"details,omitempty"`
	Result         string `json:"result,omitempty"`   // PASSED, FAILED, IGNORED, SKIPPED
	Severity       string `json:"severity,omitempty"` // HIGH, MEDIUM, LOW, CRITICAL
	Link           string `json:"link,omitempty"`
	CreatedOn      string `json:"created_on,omitempty"`
	UpdatedOn      string `json:"updated_on,omitempty"`
}
