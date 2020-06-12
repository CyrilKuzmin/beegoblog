package post

import (
	"fmt"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

//Post описывает пост
type Post struct {
	PostID      string `bson:"_id,omitempty"`
	Author		string `bson:"author,omitempty"`
	Title       string `bson:"title,omitempty"`
	ContentHTML string `bson:"contenthtml"`
	CreatedAt   string `bson:"createdat,omitempty"`
	ModifiedAt  string `bson:"modifiedat,omitempty"`
}

//NewPost создает пост и возвращает указатель на него
func NewPost(id, author, title, content string, policy *bluemonday.Policy) *Post {
	dt := time.Now().Local().Format("01-02-2006 15:04:05")
	return &Post{id, author, title, policy.Sanitize(content), dt, dt}
}

//EditPost редактируем пост
func EditPost(p *Post, title, content string, policy *bluemonday.Policy) *Post {
	p.Title = title
	p.ContentHTML = policy.Sanitize(content)
	p.ModifiedAt = time.Now().Local().Format("01-02-2006 15:04:05")
	return p
}

func (p *Post) String() string {
	return fmt.Sprintf("PostID %v %v : %v", p.PostID, p.Title, p.ContentHTML[0:40])
}
