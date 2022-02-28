package blog

import "fmt"

type Post struct {
	Id       int `json:"id"`
	AuthorId int `json:"author"`
	Author   Author
	Title    string `json:"title"`
	Body     string `json:"body"`
}

func (p Post) BasicText() string {
	return fmt.Sprintf("id: %d, author: [%s], title: %s, body: %s", p.Id, p.Author.DisplayName(), p.Title, p.Body)
}
