package controllers

import (
	"github.com/xxlaefxx/beegoblog/models/postsdb"
)

var pdb *postsdb.PostsDB

func init() {
	pdb = postsdb.NewPostsDB()
}
