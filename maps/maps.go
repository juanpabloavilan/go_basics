package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("word already exists")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}

	d[word] = definition

}

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}
