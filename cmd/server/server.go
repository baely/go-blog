package server

import (
	"encoding/json"
	"fmt"
	"go-blog/pkg/blog"
	"go-blog/pkg/util"
	"net/http"

	"github.com/go-chi/chi"
)

type DataMeta struct {
	authorsFile string
	postsFile   string
}

type Server struct {
	address string
	port    string
	blogData    util.BlogData
	meta    DataMeta
}

func (s *Server) Run(authorsFile string, postsFile string) error {
	if s.port == "" {
		s.port = "8080"
	}

	r := chi.NewRouter()

	r.Get("/authors", s.serveAuthors)
	r.Get("/posts", s.serverPosts)

	if err := s.init(); err != nil {
		return err
	}

	fmt.Println("Server is running...")

	if err := http.ListenAndServe(s.address+":"+s.port, r); err != nil {
		panic(err)
	}

	return nil
}

func (s *Server) refreshBlogData() error {
	data, err := util.LoadServerData(s.meta.authorsFile, s.meta.postsFile)
	if err != nil {
		return err
	}
	s.blogData = data

	return nil
}

func (s *Server) init() error {
	if err := s.refreshBlogData(); err != nil {
		return err
	}

	return nil
}

func (s *Server) serveAuthors(w http.ResponseWriter, r *http.Request) {
	authors := make([]blog.Author, len(s.blogData.Authors))

	i := 0
	for _, author := range s.blogData.Authors {
		authors[i] = author
		i++
	}

	authorsResponse, _ := json.Marshal(authors)
	fmt.Fprintln(w, string(authorsResponse))
}

func (s *Server) serverPosts(w http.ResponseWriter, r *http.Request) {
	posts := make([]blog.Post, len(s.blogData.Posts))

	i := 0
	for _, post := range s.blogData.Posts {
		posts[i] = post
		i++
	}

	postsResponse, _ := json.Marshal(posts)
	fmt.Fprintln(w, string(postsResponse))
}
