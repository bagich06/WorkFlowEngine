package workflow

type ApproveRequest struct {
	Comment string `json:"comment,omitempty"`
}

type RejectRequest struct {
	Reason string `json:"reason"`
}
