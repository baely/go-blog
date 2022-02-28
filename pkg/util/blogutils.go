package util

import (
	"encoding/json"
	"go-blog/pkg/blog"
	"io/ioutil"
)

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
