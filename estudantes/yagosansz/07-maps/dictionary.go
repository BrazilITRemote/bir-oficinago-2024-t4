package main

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	return d[word], nil
}
