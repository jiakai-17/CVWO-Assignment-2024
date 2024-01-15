package models

import "time"

type Comment struct {
	ID          string    `json:"id"`
	Body        string    `json:"body"`
	Creator     string    `json:"creator"`
	ThreadID    string    `json:"thread_id"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
