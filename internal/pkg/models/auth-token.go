package models

type AuthToken struct {
	Username  string `json:"username"`
	Token string `json:"token"`
}
