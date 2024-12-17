package maps_0 // Erros como constantes / Wrappers (encapsular) de erro

/*
var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)
*/
// Criando constates de Erro
// https://dave.cheney.net/2016/04/07/constant-errors

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

// Criando tipo de dados para erros
type DictionaryErr string

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

/*
func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
*/
/*
// Criando função para retorno da mensagem de erro
func (d Dictionary) Update(word, definition string) error {
	d[word] = definition
	return nil
}
*/

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

/*
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
*/

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

// Criando função para retorno da mensagem de erro
func (e DictionaryErr) Error() string {
	return string(e)
}
