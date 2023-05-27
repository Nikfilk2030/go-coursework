package user

type User struct {
	login    string
	password string
}

func NewUser(login string, password string) *User {
	return &User{login: login, password: password}
}

func (u *User) GetLogin() string {
	return u.login
}

func (u *User) GetPassword() string {
	return u.password
}
