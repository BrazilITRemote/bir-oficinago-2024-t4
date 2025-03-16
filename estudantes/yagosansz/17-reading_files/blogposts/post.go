package blogposts

import "io"

type Post struct {
	Title string
}

func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	// Here we are cutting out `Title: ` by slicing the string
	post := Post{Title: string(postData)[7:]}
	return post, nil
}
