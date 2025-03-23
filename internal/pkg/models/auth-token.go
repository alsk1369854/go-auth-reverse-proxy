package models

type AuthToken struct {
	Name  string 		`json:"name"`
	Description string 	`json:"description"`
	Token string 		`json:"token"`
}
