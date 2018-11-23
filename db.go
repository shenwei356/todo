package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

type ItemDB struct {
	db *badger.DB

	seq *badger.Sequence
}

func Connect(dir string) (*ItemDB, error) {
	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	seq, err := db.GetSequence([]byte("max_id"), 100)
	if err != nil {
		return nil, err
	}

	return &ItemDB{db: db, seq: seq}, nil
}

func (db *ItemDB) Close() error {
	err := db.seq.Release()
	if err != nil {
		return err
	}
	return db.Close()
}

func (db *ItemDB) PutItem(content []byte) (item Item, err error) {
	var i uint64
	var id []byte
	err = db.db.Update(func(txn *badger.Txn) error {
		i, err = db.seq.Next()
		if err != nil {
			return err
		}
		id = []byte(fmt.Sprintf("%d", i))
		item = Item{ID: id, Done: false, Content: content}
		ctt := MarshalItem(item)
		return txn.Set(id, ctt)
	})
	if err != nil {
		return Item{}, err
	}
	return Item{ID: id, Done: false, Content: content}, nil
}

func (db *ItemDB) GetItem(id []byte) (item Item, err error) {
	var ctt []byte
	err = db.db.View(func(txn *badger.Txn) error {
		it, err := txn.Get(id)
		if err != nil {
			return err
		}
		err = it.Value(func(val []byte) error {
			ctt = append([]byte{}, val...)
			return nil
		})
		return nil
	})
	item, err = UnmarshalItem(ctt)
	if err != nil {
		return item, err
	}
	item.ID = id
	return item, nil
}

func (db *ItemDB) GetItems(n int) (Item, error) {
	return Item{}, nil
}

func (db *ItemDB) DeleteItem(id []byte) error {
	return nil
}

func (db *ItemDB) SearchItems(query []byte) ([]Item, error) {
	return nil, nil
}
