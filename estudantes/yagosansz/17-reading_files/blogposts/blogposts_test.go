package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/BrazilITRemote/bir-oficinago-2024-t4/estudantes/yagosansz/17-reading_files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
