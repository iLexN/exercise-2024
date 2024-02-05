package model

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`

	UserID string `json:"userId"`
	User   *User  `json:"user"`
}

type User struct {
	ID   int64   `json:"id"`
	Name string  `json:"name"`
	Todo []*Todo `json:"todo,omitempty"`

	TodoIDs []string `json:"todo_ids,omitempty"`
}

type NewUser struct {
	Name string `json:"name" db:"name"`
}
