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

//ImageProcessing выдергивает base64 из поста, сохраняет в /static/img и заменяет base64 на адрес
func ImageProcessing(content string) string {
	var base64data string
	var pathToImages = "./static/img/posts/"
	result := content
	subStrArr := strings.Split(content, "\"")
	for _, elem := range subStrArr {
		//base64 картинка будет выглядеть как <img src="data:image/png;base64,iVBORw0Ktn==">
		if strings.HasPrefix(elem, "data:image") {
			base64data = strings.Split(elem, ",")[1]
			pr := strings.Split(elem, ";")[0]
			ext := strings.Split(pr, "/")[1]
			filepath := fmt.Sprintf("%v%v.%v", pathToImages, utils.GenerateUUID(), ext)
			err := base64ToFile(base64data, filepath, ext)
			if err != nil {
				fmt.Println(err)
				continue
			}
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
	id := c.Ctx.Request.FormValue("id")
	title := c.Ctx.Request.FormValue("title")
	content := c.Ctx.Request.FormValue("editor")
	if title == "" {
		//Без заголовка не принимаем
		c.Redirect("/error", 302)
		return
	}
	content = ImageProcessing(content)
	fmt.Println("Контент до политик", content)
	verifyPolicy := utils.MakeNewPolicy()
	if id != "" {
		pdb.UpdateOne(post.EditPost(pdb.SelectByID(id), title, content, verifyPolicy))
	} else {
		pdb.InsertOne(post.NewPost(utils.GenerateID(8), title, content, verifyPolicy))
	}
	c.Redirect("/blog", 302)
	c.TplName = "post.html"
	c.Layout = "layout.html"
}
