package main

import server2 "go-blog/cmd/server"

func main() {
	server := server2.Server{}
	server.Run("authors.json", "posts.json")
}
