package server

import (
	"encoding/json"
	"fmt"
	"go-blog/pkg/blog"
	"go-blog/pkg/util"
	"net/http"

	"github.com/go-chi/chi"
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

func loadServerData(authorsFile string, postsFile string) (serverData, error) {
	authors, err1 := util.LoadAuthors(authorsFile)
	if err1 != nil {
		return serverData{}, err1
	}
	posts, err2 := util.LoadPosts(postsFile, authors)
	if err2 != nil {
		return serverData{}, err2
	}
	data := serverData{
		authors: authors,
		posts:   posts,
	}
	return data, nil
}

func (s *Server) Run(authorsFile string, postsFile string) error {
	if s.port == "" {
		s.port = "8080"
	}

	r := chi.NewRouter()

	r.Get("/authors", s.serveAuthors)
	r.Get("/posts", s.serverPosts)

	if err := s.init(authorsFile, postsFile); err != nil {
		return err
	}

	fmt.Println("Server is running...")

	if err := http.ListenAndServe(s.address+":"+s.port, r); err != nil {
		panic(err)
	}

	return nil
}

func (s *Server) init(authorsFile string, postsFile string) error {
	data, err := loadServerData(authorsFile, postsFile)
	if err != nil {
		return err
	}
	s.data = data

	return nil
}

func (s *Server) serveAuthors(w http.ResponseWriter, r *http.Request) {
	authors := make([]blog.Author, len(s.data.authors))

	i := 0
	for _, author := range s.data.authors {
		authors[i] = author
		i++
	}

	authorsResponse, _ := json.Marshal(authors)
	fmt.Fprintln(w, string(authorsResponse))
}

func (s *Server) serverPosts(w http.ResponseWriter, r *http.Request) {
	posts := make([]blog.Post, len(s.data.posts))

	i := 0
	for _, post := range s.data.posts {
		posts[i] = post
		i++
	}

	postsResponse, _ := json.Marshal(posts)
	fmt.Fprintln(w, string(postsResponse))
}
