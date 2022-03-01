package blog

import "fmt"

type RawPost struct {
	Id       int    `json:"id"`
	AuthorId int    `json:"author"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

type Post struct {
	Id       int    `json:"id"`
	Author   Author `json:"author"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

func (p Post) BasicText() string {
	return fmt.Sprintf("id: %d, author: [%s], title: %s, body: %s", p.Id, p.Author.DisplayName(), p.Title, p.Body)
}

func (p RawPost) ToPost(authors map[int]Author) Post {
	return Post{
		Id: p.Id,
		Author: authors[p.AuthorId],
		Title: p.Title,
		Body: p.Body,
	}
}
