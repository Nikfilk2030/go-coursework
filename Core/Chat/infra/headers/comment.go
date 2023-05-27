package headers

type Comment interface {
	GetId() int
	GetOwnerLogin() string
	GetType() int
	GetText() string
	GetParent() int
}
