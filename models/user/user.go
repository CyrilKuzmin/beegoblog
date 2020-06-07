package user

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

//User структура описывает юзера. Содержит все его поля, включая активность и персоданные
type User struct {
	UserID       string `orm:"pk;column(user_id)"`
	UserName     string `orm:"size(20);unique"`
	Email        string `orm:"size(100);unique"`
	Password     string `orm:"size(24)"`
	PasswordHash string
	IsActive     bool
	IsBanned     bool
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	ModifiedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

//NewUser создает юзера и возвращает указатель на него
func NewUser(id, username, email, password string) *User {
	dt := time.Now().Local()
	return &User{id, username, email, password, password, true, false, dt, dt}
}

func (u User) String() string {
	return fmt.Sprintf("User %v %v Created at: %v", u.UserID, u.UserName, u.CreatedAt.Format("01-02-2006 15:04:05"))
}
