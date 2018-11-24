package main

import (
	"fmt"

	"github.com/francoispqt/gojay"
)

type Item struct {
	ID   int `storm:"id,increment"`
	Done bool
	Task string
}

func (it Item) String() string {
	return fmt.Sprintf("ID: %d, Done: %v, Task: %s", it.ID, it.Done, it.Task)
}

func (it *Item) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("id", it.ID)
	enc.BoolKey("done", it.Done)
	enc.StringKey("task", it.Task)
}

func (it *Item) IsNil() bool {
	return it == nil
}

func (it *Item) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.Int(&it.ID)
	case "done":
		return dec.Bool(&it.Done)
	case "task":
		return dec.String(&it.Task)
	}
	return nil
}

func (it *Item) NKeys() int {
	return 3
}

// ----------------------

type Items []Item

func (its *Items) MarshalJSONArray(enc *gojay.Encoder) {
	for _, it := range *its {
		enc.Object(&it)
	}
}

func (its *Items) IsNil() bool {
	return its == nil
}

func (its *Items) UnmarshalJSONArrary(dec *gojay.Decoder) error {
	it := Item{}
	if err := dec.Object(&it); err != nil {
		return err
	}
	*its = append(*its, it)
	return nil
}
