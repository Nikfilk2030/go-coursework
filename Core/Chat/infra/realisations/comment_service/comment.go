package comment_service

type Comment struct {
	Id          int    `json:"id"`
	OwnerLogin  string `json:"ownerLogin"`
	CommentType int    `json:"commentType"` // todo поправить нейминг
	Text        string `json:"text"`
	Parent      int    `json:"parent"` // todo parentId
}

func NewComment(id int, ownerLogin string, commentType int, text string, parent int) *Comment {
	return &Comment{
		Id:          id,
		OwnerLogin:  ownerLogin,
		CommentType: commentType,
		Text:        text,
		Parent:      parent,
	}
}

func (c *Comment) GetId() int {
	return c.Id
}
func (c *Comment) GetOwnerLogin() string {
	return c.OwnerLogin
}
func (c *Comment) GetType() int {
	return c.CommentType
}
func (c *Comment) GetText() string {
	return c.Text
}
func (c *Comment) GetParent() int {
	return c.Parent
}
