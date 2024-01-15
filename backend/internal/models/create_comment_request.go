package models

// CreateCommentRequest Provides the layout for the JSON object sent by frontend to create a comment
type CreateCommentRequest struct {
	ThreadId string `json:"thread_id"`
	Body     string `json:"body"`
}
