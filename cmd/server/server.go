package server

import (
	"fmt"
	"go-blog/pkg/blog"
	"net/http"
)

type serverData struct {
	authors map[int]blog.Author
	posts   map[int]blog.Post
}

type Server struct {
	address string
	port    string
	data    serverData
}

func loadServerData(authorsFile string, postsFile string) serverData {
	authors := blog.LoadAuthors(authorsFile)
	posts := blog.LoadPosts(postsFile, authors)
	data := serverData{
		authors: authors,
		posts:   posts,
	}
	return data
}

func (s *Server) serveAuthors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serverAuthors was hit")
	fmt.Println("length of authors is: ", s.data.authors)

	for i, author := range s.data.authors {
		if _, err := fmt.Fprintf(w, "Author #%d: [%s]\n", i, author.DisplayName()); err != nil {
			fmt.Println(err)
		}
	}
}

func (s *Server) serverPosts(w http.ResponseWriter, r *http.Request) {
	for i, post := range s.data.posts {
		if _, err := fmt.Fprintf(w, "Post #%d: [%s]\n", i, post.BasicText()); err != nil {
			fmt.Println(err)
		}
	}
}

func (s *Server) Run(authorsFile string, postsFile string) {
	if s.address == "" {
		s.address = "0.0.0.0"
	}
	if s.port == "" {
		s.port = "8080"
	}

	s.data = loadServerData(authorsFile, postsFile)

	http.HandleFunc("/authors", s.serveAuthors)
	http.HandleFunc("/posts", s.serverPosts)

	fmt.Println("Server is running...")

	if err := http.ListenAndServe(s.address+":"+s.port, nil); err != nil {
		fmt.Println(err)
	}
}
