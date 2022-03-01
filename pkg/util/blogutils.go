package util

import (
	"encoding/json"
	"go-blog/pkg/blog"
	"io/ioutil"
	"os"
)

type BlogData struct {
	Authors map[int]blog.Author
	Posts   map[int]blog.Post
}

func LoadServerData(authorsFile string, postsFile string) (BlogData, error) {
	authors, err1 := LoadAuthors(authorsFile)
	if err1 != nil {
		return BlogData{}, err1
	}
	posts, err2 := LoadPosts(postsFile, authors)
	if err2 != nil {
		return BlogData{}, err2
	}
	data := BlogData{
		Authors: authors,
		Posts:   posts,
	}
	return data, nil
}

func LoadAuthors(authorsLocation string) (map[int]blog.Author, error) {
	var authors []blog.Author
	authorsMap := make(map[int]blog.Author)

	authorData, err := ioutil.ReadFile(authorsLocation)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(authorData, &authors)
	if err != nil {
		return nil, err
	}

	for _, author := range authors {
		authorsMap[author.Id] = author
	}

	return authorsMap, err
}

func SaveAuthor(authorsLocation string, author blog.Author) error {
	authorsMap, err := LoadAuthors(authorsLocation)
	if err != nil {
		return err
	}
	authors := make([]blog.Author, len(authorsMap))

	i := 0
	for _, existingAuthor := range authorsMap {
		authors[i] = existingAuthor
		i++
	}

	author.Id = i

	authors = append(authors, author)
	authorsData, err := json.Marshal(authors)
	if err != nil {
		return err
	}

	os.WriteFile(authorsLocation, authorsData, 0777)

	return nil
}

func LoadPosts(postsLocation string, authors map[int]blog.Author) (map[int]blog.Post, error) {
	var posts []blog.RawPost
	postsMap := make(map[int]blog.Post)

	postData, err := ioutil.ReadFile(postsLocation)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(postData, &posts)
	if err != nil {
		return nil, err
	}

	var post blog.Post
	for _, rawPost := range posts {
		post = rawPost.ToPost(authors)
		postsMap[post.Id] = post
	}

	return postsMap, nil
}
