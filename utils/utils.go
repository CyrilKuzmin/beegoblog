package utils

import (
	"html/template"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//GenerateUUID генерирует UUID (128 bits)
func GenerateUUID() string {
	return uuid.New().String()
}

//StringWithCharset генерирует строку заданной длины, используя список символов charset
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//GenerateID генерирует UUID (16-20 символов)
func GenerateID(length int) string {
	return StringWithCharset(length, charset)
}

//Unescape используется для отображения поста в красивом HTML вместо чистого кода HTML
func Unescape(x string) interface{} {
	return template.HTML(x)
}

//MDToHTML принимает на вход строку формата MD и возвращает кусок HTML (строку)
func MDToHTML(md string) string {
	htmlBytes := blackfriday.MarkdownBasic([]byte(md))
	return string(htmlBytes)
}

//MakeNewPolicy создает новую политику bluemonday для верификации пользовательских данных, полученных на метод /savepost
func MakeNewPolicy() *bluemonday.Policy {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("style").OnElements("span", "p")
	return p
}
