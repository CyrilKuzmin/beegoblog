package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/xxlaefxx/beegoblog/models/user"
	"github.com/xxlaefxx/beegoblog/utils"
)

//RegisterController служит для регистрации новых пользователей
type RegisterController struct {
	beego.Controller
}

//RegistrationError модель для сообщений при регистрации
type RegistrationMessage struct {
	Label    string
	LinkPath string
	LinkText string
}

//NewRegistrationMessage Используется для создания сообщений о регистрации
func NewRegistrationMessage(label, linkPath, linktext string) *RegistrationMessage {
	return &RegistrationMessage{label, linkPath, linktext}
}

/*Эти значения используются для информацирования юзера о (без)успешности реги
Сообщение содержит label и ссылку.
*/
var userAlreadyExists = NewRegistrationMessage("Пользователь с таким именем уже существует. ", "/register", " Повторить попытку")
var emailAlreadyExists = NewRegistrationMessage("Пользователь с таким email уже существует. ", "/register", " Повторить попытку")
var passwordWrong = NewRegistrationMessage(`Нельзя использовать данный пароль. 
Он может быть длиной от 8 до 20 символов и состоять из латинских букв, цифр и символов: 
\@ \$ \! \% \* \? \& `, "/register", " Повторить попытку")
var everythingIsOk = NewRegistrationMessage("Регистрация прошла успешно. Теперь вы можете ", "/login", " войти")

func userExists(username string) bool {
	u := user.User{UserName: username}
	err := psqlOrm.Read(&u, "user_name")
	if err != orm.ErrNoRows {
		return true
	}
	return false
}

func emailExists(email string) bool {
	u := user.User{Email: email}
	err := psqlOrm.Read(&u, "email")
	if err != orm.ErrNoRows {
		return true
	}
	return false
}

func wrongPassword(password string) bool {
	return false
}

func runChecks(username, email, password string) *RegistrationMessage {
	if userExists(username) {
		return userAlreadyExists
	}
	if emailExists(email) {
		return emailAlreadyExists
	}
	if wrongPassword(password) {
		return passwordWrong
	}
	return everythingIsOk
}

//Get возвращает страничку регистрации
func (c *RegisterController) Get() {
	c.TplName = "register.html"
	c.Layout = "layout.html"
}

//Post проверяет данные регистрации и регистрирует пользователя
func (c *RegisterController) Post() {
	username := c.Ctx.Request.FormValue("username")
	email := c.Ctx.Request.FormValue("email")
	password := c.Ctx.Request.FormValue("password")
	//Проверки
	checksResult := runChecks(username, email, password)
	c.Data["Message"] = checksResult
	if checksResult == everythingIsOk {
		_, err := psqlOrm.Insert(user.NewUser(utils.GenerateUUID(), username, email, password))
		if err != nil {
			fmt.Println(err)
		}
	}
	c.TplName = "register.html"
	c.Layout = "layout.html"
}
