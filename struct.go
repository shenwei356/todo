package main

import (
	"fmt"
)

type Item struct {
	ID      int `storm:"id,increment"`
	Done    bool
	Content string
}

func (it Item) String() string {
	return fmt.Sprintf("ID: %d, Done: %v, Content: %s", it.ID, it.Done, it.Content)
}
