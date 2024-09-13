package model

type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JWTResponse struct {
	JWT string `json:"jwt"`
}
