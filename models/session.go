package models

type Session struct {
	Id     string `json:"id"`
	User   string `json:"user"`
	Email  string `json:"email"`
	Valid  string `json:"valid"`
	Admin  string `json:"admin"`
	Login  string `json:"login"`
	TeamId string `json:"team_id"`
	Locale string `json:"locale"`
}
