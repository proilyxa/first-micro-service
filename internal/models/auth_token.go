package models

type AuthToken struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"userId"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}
