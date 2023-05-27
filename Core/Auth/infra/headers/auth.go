package headers

type Auth interface {
	Login(login string, password string) error
	Auth() error
	Register(login string, password string) error
}
