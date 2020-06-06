package postdocuments

import (
	"context"
	"fmt"
	"time"

	"github.com/xxlaefxx/beegoblog/models/post"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//PostDocuments объект для взаимодействия с MongoDB
type PostDocuments struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

//NewPostDocuments создает объект для взаимодействия с MongoDB
func NewPostDocuments() *PostDocuments {
	var mongoURI = "mongodb://localhost:27017"
	var dbName = "blog"
	var collectionName = "posts"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	ctx, mongoConnectCancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer mongoConnectCancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return &PostDocuments{client, db, collection}
}

//InsertOne инсертит 1 пост в коллекцию
func (pD PostDocuments) InsertOne(post *post.Post) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	pD.collection.InsertOne(ctx, post)
}

//UpdateOne инсертит 1 пост в коллекцию
func (pD PostDocuments) UpdateOne(post *post.Post) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	pD.collection.FindOneAndReplace(ctx, bson.M{"_id": bson.M{"$eq": post.ID}}, post)
}

//DeleteByID удаляет 1 пост в коллекцию
func (pD PostDocuments) DeleteByID(id string) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	_, err := pD.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		fmt.Printf("Cannot delete post %v, %v\n", id, err)
	}
}

//SelectAll возвращает все посты
func (pD PostDocuments) SelectAll() *[]post.Post {
	var posts []post.Post
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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
func (pD PostDocuments) SelectByQuery(query bson.M) *[]post.Post {
	var posts []post.Post
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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
func (pD PostDocuments) SelectByID(id string) *post.Post {
	var post post.Post
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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
