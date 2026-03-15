package feature2

import "time"

type Data struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
