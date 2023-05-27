package headers

type Storage interface {
	GetComments(parentId int) (string, error)
}
