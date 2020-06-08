package controllers

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/xxlaefxx/beegoblog/models/user"
	"github.com/xxlaefxx/beegoblog/utils"
)

//RegisterController служит для регистрации новых пользователей
type RegisterController struct {
	beego.Controller
}

//RegistrationMessage модель для сообщений при регистрации
type RegistrationMessage struct {
	Label    string
	LinkPath string
	LinkText string
	Level    string
}

//NewRegistrationMessage Используется для создания сообщений о регистрации
func NewRegistrationMessage(label, linkPath, linktext, level string) *RegistrationMessage {
	return &RegistrationMessage{label, linkPath, linktext, level}
}

/*Эти значения используются для информацирования юзера о (без)успешности реги
Сообщение содержит label и ссылку.
*/
var userAlreadyExists = NewRegistrationMessage("Пользователь с таким именем уже существует. ", "", "", "warning")
var emailAlreadyExists = NewRegistrationMessage("Пользователь с таким email уже существует. ", "", "", "warning")
var passwordWrong = NewRegistrationMessage(`Нельзя использовать данный пароль. 
Он может быть длиной от 8 до 20 символов и состоять из латинских букв, цифр и символов: 
@ $ ! % * ? & `, "", "", "warning")
var everythingIsOk = NewRegistrationMessage("Регистрация прошла успешно. Теперь вы можете ", "/login", " войти", "success")
var passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d@$!%*?&]{8,20}$`

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
	//RE2 не позволит взять и сделать красиво, как в JS/Python...
	if len(password) < 8 {
		return true
	}
	if strings.Contains(password, " ") {
		return true
	}
	noNumbers := true
	noUppers := true
	noLowers := true
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			noNumbers = false
		case unicode.IsUpper(c):
			noUppers = false
		case unicode.IsLetter(c):
			noLowers = false
		default:
		}
	}
	result := noNumbers || noUppers || noLowers
	return result
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
	sess := c.GetSession("session")
	if sess != nil {
		m := sess.((map[string]interface{}))
		c.Data["UserName"] = m["username"]
	} else {
		c.Abort("403")
	}
	c.TplName = "register.html"
	c.Layout = "layout.html"
}

//Post проверяет данные регистрации и регистрирует пользователя
func (c *RegisterController) Post() {
	fmt.Println(everythingIsOk)
	username := c.Ctx.Request.FormValue("username")
	email := c.Ctx.Request.FormValue("email")
	password := c.Ctx.Request.FormValue("password")
	//Проверки
	checksResult := runChecks(username, email, password)
	fmt.Println(checksResult)
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
