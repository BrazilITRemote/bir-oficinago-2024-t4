package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// Map look up returns 2 values.
	// The 2nd one being a boolean to indicate if the key was found!
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
