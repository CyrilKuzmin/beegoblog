package postsdb

import (
	"context"
	"fmt"
	"time"

	"github.com/xxlaefxx/beegoblog/models/post"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//PostsDB объект для взаимодействия с MongoDB
type PostsDB struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
	bgCtx      context.Context
}

//NewPostsDB создает объект для взаимодействия с MongoDB
func NewPostsDB() *PostsDB {
	var mongoURI = "mongodb://localhost:27017"
	var dbName = "blog"
	var collectionName = "posts"
	var bgCtx = context.TODO()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(bgCtx, 20*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return &PostsDB{client, db, collection, bgCtx}
}

//InsertOne инсертит 1 пост в коллекцию
func (pD PostsDB) InsertOne(post *post.Post) {
	ctx, cancel := context.WithTimeout(pD.bgCtx, 20*time.Second)
	defer cancel()
	pD.collection.InsertOne(ctx, post)
}

//UpdateOne инсертит 1 пост в коллекцию
func (pD PostsDB) UpdateOne(post *post.Post) {
	ctx, cancel := context.WithTimeout(pD.bgCtx, 20*time.Second)
	defer cancel()
	pD.collection.FindOneAndReplace(ctx, bson.M{"_id": bson.M{"$eq": post.PostID}}, post)
}

//DeleteByID удаляет 1 пост в коллекцию
func (pD PostsDB) DeleteByID(id string) {
	ctx, cancel := context.WithTimeout(pD.bgCtx, 20*time.Second)
	defer cancel()
	_, err := pD.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		fmt.Printf("Cannot delete post %v, %v\n", id, err)
	}
}

//SelectAll возвращает все посты
func (pD PostsDB) SelectAll() *[]post.Post {
	var posts []post.Post
	ctx, cancel := context.WithTimeout(pD.bgCtx, 20*time.Second)
	defer cancel()
	postsCursor, err := pD.collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Cannot find docs: %v\n", err)
	}
	for postsCursor.Next(ctx) {
		var elem post.Post
		err := postsCursor.Decode(&elem)
		if err != nil {
			fmt.Printf("Cannot decode docs: %v\n", err)
		}
		posts = append(posts, elem)
	}
	return &posts
}

//SelectByQuery возвращает документы по данному запросу (query)
func (pD PostsDB) SelectByQuery(query bson.M) *[]post.Post {
	var posts []post.Post
	ctx, cancel := context.WithTimeout(pD.bgCtx, 20*time.Second)
	defer cancel()
	postsCursor, err := pD.collection.Find(ctx, query)
	if err != nil {
		fmt.Printf("Cannot find docs: %v\n", err)
	}
	for postsCursor.Next(ctx) {
		var elem post.Post
		err := postsCursor.Decode(&elem)
		if err != nil {
			fmt.Printf("Cannot decode docs: %v\n", err)
		}
		posts = append(posts, elem)
	}
	return &posts
}

//SelectByID возвращает один пост по его ID
func (pD PostsDB) SelectByID(id string) *post.Post {
	var post post.Post
	ctx, cancel := context.WithTimeout(pD.bgCtx, 20*time.Second)
	defer cancel()
	postsCursor, err := pD.collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		fmt.Printf("Cannot find a doc: %v\n", err)
	}
	for postsCursor.Next(ctx) {
		err := postsCursor.Decode(&post)
		if err != nil {
			fmt.Printf("Cannot decode a doc: %v\n", err)
		}
	}
	return &post
}
