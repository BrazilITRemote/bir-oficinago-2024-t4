package main

import (
	"log"
	"os"

	"github.com/BrazilITRemote/bir-oficinago-2024-t4/estudantes/yagosansz/17-reading_files/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
