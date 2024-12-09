package api

type RunRequest struct {
	Slug       string `json:"slug"`
	Submission string `json:"submission"`
}

type RunResponse struct {
	State    string `json:"state"`
	FailAtID int    `json:"failed_at_id"`
	Message  string `json:"message,omitempty"`
}
