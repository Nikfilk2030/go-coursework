package headers

type Comments interface {
	GetComments(parentId int) (string, error)
}
