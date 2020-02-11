package model

type TSysUser struct {
	UserId      string
	Password    string
	Email       string
	PhoneNumber string
	Dept        string
	Name        string
	UserAccount string
}

type JwtToken struct {
	Token string `json:"token"`
}
