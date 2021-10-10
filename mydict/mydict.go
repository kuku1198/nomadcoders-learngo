package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCanUpdate  = errors.New("Cant update non-existion word")
	errWordExists = errors.New("That word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCanUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}