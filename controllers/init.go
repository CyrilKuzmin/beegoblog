package controllers

import (
	"fmt"
	"os/user"

	"github.com/astaxie/beego/orm"
	"github.com/xxlaefxx/beegoblog/models/postsdb"
)

var pdb *postsdb.PostsDB
var psqlOrm orm.Ormer

func init() {
	pdb = postsdb.NewPostsDB()
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=root password=root host=127.0.0.1 port=5432 dbname=blog sslmode=disable")
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating user model for ORM", new(user.User))
	psqlOrm = orm.NewOrm()
	psqlOrm.Using("default")
}
