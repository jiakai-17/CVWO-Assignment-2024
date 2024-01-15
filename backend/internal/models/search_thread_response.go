package models

type SearchThreadResponse struct {
	TotalThreads int32    `json:"total_threads"`
	Threads      []Thread `json:"threads"`
}
