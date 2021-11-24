package models

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ActiveToken struct {
	Token string `json:"token"`
}

type ErrorDescr struct {
	Reason string `json:"reason"`
}
