package blogrenderer_test

import (
	"bytes"
	"testing"

	blogrenderer "github.com/BrazilITRemote/bir-oficinago-2024-t4/estudantes/yagosansz/18-templating"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>`

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
