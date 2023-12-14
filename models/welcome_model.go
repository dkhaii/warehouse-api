package models

type GetWelcomeMesssage struct {
	AppName   string `json:"app_name"`
	Developer string `json:"developer"`
	Message   string `json:"message"`
}
