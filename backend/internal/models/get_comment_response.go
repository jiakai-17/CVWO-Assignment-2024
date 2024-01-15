package models

type GetCommentResponse struct {
	Comments []Comment `json:"comments"`
	Count    int32     `json:"count"`
}
