// package repository

// import (
// 	"context"
// 	"log"

// 	"cloud.google.com/go/firestore"
// 	"github.com/Delaram-Gholampoor-Sagha/Golang_Redis/entity"

// 	"google.golang.org/api/iterator"
// )

// type repo struct {
// 	CollectionName string
// }

// //NewFirestoreRepository creates a new repo
// func NewFirestoreRepository(collectionName string) PostRepository {
// 	return &repo{
// 		CollectionName: collectionName,
// 	}
// }

// const (
// 	projectID string = "pragmatic-reviews"
// )

// func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, projectID)
// 	if err != nil {
// 		log.Fatalf("Failed to create a Firestore Client: %v", err)
// 		return nil, err
// 	}

// 	defer client.Close()
// 	_, _, err = client.Collection(r.CollectionName).Add(ctx, map[string]interface{}{
// 		"ID":    post.ID,
// 		"Title": post.Title,
// 		"Text":  post.Text,
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed adding a new post: %v", err)
// 		return nil, err
// 	}
// 	return post, nil
// }

// func (r *repo) FindAll() ([]entity.Post, error) {
// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, projectID)
// 	if err != nil {
// 		log.Fatalf("Failed to create a Firestore Client: %v", err)
// 		return nil, err
// 	}

// 	defer client.Close()
// 	var posts []entity.Post
// 	it := client.Collection(r.CollectionName).Documents(ctx)
// 	for {
// 		doc, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Failed to iterate the list of posts: %v", err)
// 			return nil, err
// 		}
// 		post := entity.Post{
// 			ID:    doc.Data()["ID"].(int64),
// 			Title: doc.Data()["Title"].(string),
// 			Text:  doc.Data()["Text"].(string),
// 		}
// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }

// //FindByID: TODO
// func (r *repo) FindByID(id string) (*entity.Post, error) {
// 	return nil, nil
// }

// //Delete: TODO
// func (r *repo) Delete(post *entity.Post) error {
// 	return nil
// }
