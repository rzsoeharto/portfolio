package models

type User struct {
	UserID    int    `json:"userid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}
