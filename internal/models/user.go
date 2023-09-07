package models

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Token     string `json:"-"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}
