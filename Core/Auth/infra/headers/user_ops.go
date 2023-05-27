package headers

type Store interface {
	GetUser(login string, password string) (User, error)
	FindUser(login string) (bool, error)
	InsertUser(user User) error
}
