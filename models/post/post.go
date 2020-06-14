package post

import (
	"fmt"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

//Post описывает пост
type Post struct {
	PostID      string `bson:"_id,omitempty"`
	Author      string `bson:"author,omitempty"`
	Title       string `bson:"title,omitempty"`
	ContentHTML string `bson:"contenthtml"`
	PreviewHTML string `bson:"previewhtml"`
	CreatedAt   string `bson:"createdat,omitempty"`
	ModifiedAt  string `bson:"modifiedat,omitempty"`
}

//NewPost создает пост и возвращает указатель на него
func NewPost(id, author, title, content string, policy *bluemonday.Policy) *Post {
	dt := time.Now().Local().Format("01-02-2006 15:04:05")
	verifiedContent := policy.Sanitize(content)
	previewContent := makePreview(verifiedContent, id)

	return &Post{id, author, title, verifiedContent, previewContent, dt, dt}
}

//EditPost редактируем пост
func EditPost(p *Post, title, content string, policy *bluemonday.Policy) *Post {
	p.Title = title
	p.ContentHTML = policy.Sanitize(content)
	p.PreviewHTML = makePreview(p.ContentHTML, p.PostID)
	p.ModifiedAt = time.Now().Local().Format("01-02-2006 15:04:05")
	return p
}

func (p *Post) String() string {
	return fmt.Sprintf("PostID %v %v : %v", p.PostID, p.Title, p.ContentHTML[0:40])
}

// makePreview возвращает начало поста, что в итоге будем отображать в блоге.
// для просмотра полного текста нужно будет переходить
func makePreview(content, id string) string {
	//если пост длинной менее 500 символов, просто вернем его, нечего обрезать
	if len(content) < 500 {
		return content
	}
	//если есть двойной <br>, это и будет разделением
	breaksIndex := strings.Index(content, "<br><br>")
	if breaksIndex > 0 {
		return addReadMore(content[0:breaksIndex], id)
	}
	//Иначе ищем конец любого тэга после 500-го символа и возвращаем
	lastTagIndex := strings.Index(content[500:], "</")
	tail := strings.Index(content[lastTagIndex:], ">")
	lastTagIndex += tail + 1
	return addReadMore(content[0:lastTagIndex], id)
}

// addReadMore добавляет ссылку на пост с текстом "Читать далее" в конец превью
func addReadMore(text, id string) string {
	return fmt.Sprintf("%v<br><a href=\"post?id=%v\">Читать далее</a>", text, id)
}
