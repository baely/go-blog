package main

import (
	"fmt"
	server2 "go-blog/cmd/server"
)

func main() {
	server := server2.Server{}
	if err := server.Run("authors.json", "posts.json"); err != nil {
		fmt.Println(err)
	}
}
