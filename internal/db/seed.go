package db

import (
	"context"
	"encoding/json"
	"io"
	"math/rand"

	"log"
	"os"

	"github.com/ivanpaghubasan/go-social/internal/store"
)

func Seed(store store.Storage) {
	ctx := context.Background()
	users := generateUsers()
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user: ", err)
			return
		}
	}

	posts := generatePosts(users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	comments := generateComments(users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error create comment: ", err)
			return
		}
	}

	log.Println("Seeding Completed!")

}

func generateUsers() []*store.User {
	usersFile, err := os.Open("mocks/users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer usersFile.Close()

	usersJson, err := io.ReadAll(usersFile)
	if err != nil {
		log.Fatal(err)
	}

	var users []*store.User
	json.Unmarshal(usersJson, &users)

	return users
}

func generatePosts(users []*store.User) []*store.Post {
	postsFile, err := os.Open("mocks/posts.json")
	if err != nil {
		log.Fatal(err)
	}
	defer postsFile.Close()

	postsJson, err := io.ReadAll(postsFile)
	if err != nil {
		log.Fatal(err)
	}

	var posts []*store.Post
	json.Unmarshal(postsJson, &posts)

	for _, post := range posts {
		user := users[rand.Intn(len(users))]
		post.UserID = user.ID
	}

	return posts
}

func generateComments(users []*store.User, posts []*store.Post) []*store.Comment {
	commentsFile, err := os.Open("mocks/comments.json")
	if err != nil {
		log.Fatal(err)
	}
	defer commentsFile.Close()

	commentsJson, err := io.ReadAll(commentsFile)
	if err != nil {
		log.Fatal(err)
	}

	var comments []*store.Comment
	json.Unmarshal([]byte(commentsJson), &comments)

	for _, comm := range comments {
		comm.UserID = users[rand.Intn(len(users))].ID
		comm.PostID = posts[rand.Intn(len(posts))].ID
	}

	return comments

}
