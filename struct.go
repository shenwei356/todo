package main

import (
	"errors"
	"fmt"
)

type Item struct {
	ID      []byte
	Done    bool
	Content []byte
}

func (it Item) String() string {
	return fmt.Sprintf("ID: %s, Done: %v, Content: %s", it.ID, it.Done, it.Content)
}

func MarshalItem(it Item) []byte {
	data := make([]byte, 1, 2)
	if it.Done {
		data[0] = 1
	}
	data = append(data, it.Content...)
	return data
}

var ErrInvalidItem = errors.New("db: invalid marshalled item data")

func UnmarshalItem(val []byte) (Item, error) {
	switch len(val) {
	case 0:
		return Item{}, ErrInvalidItem
	case 1:
		return Item{Done: val[0] != 0, Content: []byte{}}, nil
	default:
	}
	return Item{Done: val[0] != 0, Content: val[1:]}, nil
}
