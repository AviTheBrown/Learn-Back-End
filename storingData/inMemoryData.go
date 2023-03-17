package main

import (
	"fmt"
)

// creates Post struct
type Post struct {
	Id      int
	Content string
	Author  string
}

// PostById -> creates a map that will be able to identiy and Post's Id
// byt looking up a  number and them locating that post
// by ppinting to the Post
var PostById map[int]*Post

// PostbyAuthor -> creates a map that links a stirng to
// an array of Post intances.
var PostByAuthor map[string][]*Post

func storeData(post *Post) {
	// PostById -> this creates a map between the post.Id from the post passed in as an argument
	// and an actual post instance.
	PostById[post.Id] = post
	// this creates a map between the post.Author and a post instance passed in
	// both of these things will enable the lookup of contents from a post instance
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author],post)
}

func main() {
	// PostbyId -> creates an instance of the PostById object create above.
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{
		Id:      1,
		Content: "i love r/Golang",
		Author:  "Andy",
	}
	post2 := Post{
		Id:      2,
		Content: "Hello Go",
		Author:  "Mary",
	}
	post3 := Post{
		Id:      3,
		Content: "Why am i here",
        // whitespaces do mater in string
		// is there is a space before or after the name
		// it will not correspond to a map witht the same name
		// without any spaces
		Author:  "Andy",
	}

	storeData(&post1)
	storeData(&post2)
	storeData(&post3)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])
	fmt.Println(PostById[3])

	// this will loop over the PostByUthor map
	// find all instances of post with the of "Andy" as the authour
	for _, post := range PostByAuthor["Andy"] {
		fmt.Println(post.Content)
	}
	for _, post := range PostByAuthor["Mary"] {
		fmt.Println(post.Content)
	}

}
