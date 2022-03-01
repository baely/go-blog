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
	if err := s.init(authorsFile, postsFile); err != nil {
		return err
	}

	r := chi.NewRouter()

	r.Get("/authors", s.getAuthors)
	r.Get("/posts", s.getPosts)

	r.Post("/authors", s.postAuthors)
	r.Post("/posts", s.postPosts)



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

func (s *Server) init(authorsFile string, postsFile string) error {
	if s.port == "" {
		s.port = "8080"
	}

	s.meta.authorsFile = authorsFile
	s.meta.postsFile = postsFile

	if err := s.refreshBlogData(); err != nil {
		return err
	}

	return nil
}

func (s *Server) getAuthors(w http.ResponseWriter, r *http.Request) {
	authors := make([]blog.Author, len(s.blogData.Authors))

	i := 0
	for _, author := range s.blogData.Authors {
		authors[i] = author
		i++
	}

	authorsResponse, _ := json.Marshal(authors)
	fmt.Fprintln(w, string(authorsResponse))
}

func (s *Server) postAuthors(w http.ResponseWriter, r *http.Request) {
	var author blog.Author

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&author)

	if err != nil {
		fmt.Fprintln(w, "Failed")
	} else {
		util.SaveAuthor(s.meta.authorsFile, author)
	}

	s.refreshBlogData()
}

func (s *Server) getPosts(w http.ResponseWriter, r *http.Request) {
	posts := make([]blog.Post, len(s.blogData.Posts))

	i := 0
	for _, post := range s.blogData.Posts {
		posts[i] = post
		i++
	}

	postsResponse, _ := json.Marshal(posts)
	fmt.Fprintln(w, string(postsResponse))
}

func (s *Server) postPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not working at this stage.")
}
