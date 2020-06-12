package controllers

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/models/post"
	"github.com/xxlaefxx/beegoblog/utils"
)

//SavePostController сохраняет пост
type SavePostController struct {
	beego.Controller
}

func base64ToFile(base64data, filepath, ext string) error {
	var err error
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64data))
	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		return err
	}
	//Пора работаем только с PNG и JPG
	switch ext {
	case "png":
		err = png.Encode(f, img)
	case "jpeg":
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	default:
		err = fmt.Errorf("Cannot work with given file format: %v", ext)
	}
	if err != nil {
		return err
	}
	return nil
}

// ImageProcessing выдергивает base64 из поста, сохраняет в /static/img
// и заменяет base64 на адрес
//
// В качестве имени файла использую UUID
//
// Директория для всех картинок: "./static/img/posts/"
func ImageProcessing(content string) string {
	//base64 картинка будет выглядеть как <img src="data:image/png;base64,iVBORw0Ktn==">
	//В эту переменную будем записывать данные в base64 (iVBORw0Ktn==)
	var base64data string
	var pathToImages = "./static/img/posts/"
	//вернем как минимум то, что получили
	result := content
	//Разделяем исходный текст на слайс по кавычке и идем по каждому элементу (elem)
	// elem[0] = "<img src="   elem[1] = "data:image/png;base64,iVBORw0Ktn==" elem[2] = ">"
	subStrArr := strings.Split(content, "\"")
	for _, elem := range subStrArr {
		//первое условие - элемент должен начинаться с "data:image"
		if strings.HasPrefix(elem, "data:image") {
			//если это так, то вычленяем base64, формат файла (ext), а также формируем имя
			base64data = strings.Split(elem, ",")[1]
			pr := strings.Split(elem, ";")[0]
			ext := strings.Split(pr, "/")[1]
			filepath := fmt.Sprintf("%v%v.%v", pathToImages, utils.GenerateUUID(), ext)
			//вызываем функцию для декодирования и сохранения
			err := base64ToFile(base64data, filepath, ext)
			if err != nil {
				fmt.Println(err)
				//Если что-то не так, переходим к следующему элементу, оставляя все как есть
				continue
			}
			//заменяем элемент "data:image/png;base64,iVBORw0Ktn==" на путь к файлу
			result = strings.ReplaceAll(result, elem, filepath)
		}
	}
	return result
}

//Post сохраняет пост
func (c *SavePostController) Post() {
	//Только зарегистрированные могут что-то постить
	if c.Data["UserName"] == nil {
		c.Abort("403")
	}
	//Берем данные с HTML-формы
	id := c.Ctx.Request.FormValue("id")
	title := c.Ctx.Request.FormValue("title")
	content := c.Ctx.Request.FormValue("editor")
	if title == "" {
		//Без заголовка не принимаем
		c.Redirect("/error", 302)
		return
	}
	//Редактор загружает картинки в Base64. Нужно схоронить на диск и заменить base64 ссылкой
	content = ImageProcessing(content)
	//Дабы нам не ввели явоскрипт и прочие не секурные вещи в тело поста:
	verifyPolicy := utils.MakeNewPolicy()
	var p *post.Post
	if id != "" {
		//Обновляем существующий пост
		p = post.EditPost(pdb.SelectByID(id), title, content, verifyPolicy)
		pdb.UpdateOne(p)
	} else {
		//Создаем новый
		p = post.NewPost(utils.GenerateID(8), c.Data["UserName"].(string), title, content, verifyPolicy)
		pdb.InsertOne(p)
	}
	//Возвращаем на результат
	c.Redirect("/post?id="+p.PostID, 302)
}
