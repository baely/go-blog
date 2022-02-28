package blog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	Id       int `json:"id"`
	AuthorId int `json:"author"`
	Author   Author
	Title    string `json:"title"`
	Body     string `json:"body"`
}

func (a *Author) DisplayName() string {
	return fmt.Sprintf("id: %d, name: %s", a.Id, a.Name)
}

func (p *Post) BasicText() string {
	return fmt.Sprintf("id: %d, author: [%s], title: %s, body: %s", p.Id, p.Author.DisplayName(), p.Title, p.Body)
}

func LoadAuthors(authorsLocation string) map[int]Author {
	var authors []Author
	authorsMap := make(map[int]Author)

	authorData, err := ioutil.ReadFile(authorsLocation)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(authorData, &authors)
	if err != nil {
		fmt.Println(err)
	}

	for _, author := range authors {
		authorsMap[author.Id] = author
	}

	return authorsMap
}

func LoadPosts(postLocation string, authors map[int]Author) map[int]Post {
	var posts []Post
	postsMap := make(map[int]Post)

	postData, err := ioutil.ReadFile(postLocation)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(postData, &posts)
	if err != nil {
		fmt.Println(err)
	}

	for _, post := range posts {
		post.Author = authors[post.AuthorId]
		postsMap[post.Id] = post
	}

	return postsMap
}
