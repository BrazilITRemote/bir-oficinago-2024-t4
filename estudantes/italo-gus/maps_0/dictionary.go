package maps_0 // Refactor - Refatoração

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

/*
func (d Dictionary) Search(word string) string {
	return d[word]
}
*/

// função com retorno de tratamento de erro
/*
func (d Dictionary) Search(word string) (string, error) {
	return d[word], nil
}
*/

/*
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
*/

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
