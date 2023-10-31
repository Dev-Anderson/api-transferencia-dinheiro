package schemas

import "time"

type UserCreate struct {
	Name     string  `json:"name"`
	Document string  `json:"document"`
	Balance  float64 `json:"balance"`
}

type User struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Document         string    `json:"document"`
	DateCreate       time.Time `json:"datecreate"`
	Balance          float64   `json:"balance"`
	DateAtualization time.Time `json:"dateatualization"`
}
