package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xxlaefxx/beegoblog/models/postsdb"
	"github.com/xxlaefxx/beegoblog/models/user"
)

var pdb *postsdb.PostsDB
var psqlOrm orm.Ormer

func init() {
	var err error
	pdb, err = postsdb.NewPostsDB("mongodb://172.16.1.1:27017")
	if err != nil {
		logs.Critical("Cannot create Mongo client: %v", err)
	}
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=root password=root host=172.16.1.1 port=5432 dbname=blog sslmode=disable")
	_ = orm.RunSyncdb("default", false, true)
	logs.Info("Creating user model for ORM", new(user.User))
	psqlOrm = orm.NewOrm()
	psqlOrm.Using("default")
}
