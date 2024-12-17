package maps_0 // Função de cadastro no mapa, aprimorada para verificar se a chave (índice) já é cadastrada

import "errors"

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

// função com retorno de tratamento de erro
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

/*
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
*/

/*
// função com retorno de tratamento de erro
func (d Dictionary) Add(word, definition string) error {
	d[word] = definition
	return nil
}
*/

// função com retorno de tratamento de erro
// função aprimorada para verificar se a chave (índice) já é cadastrada
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
