package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/asdine/storm"
)

type ItemDB struct {
	db *storm.DB
}

func Connect(path string) (*ItemDB, error) {
	dir := filepath.Dir(path)
	existed, err := DirExists(dir)
	if err != nil {
		log.Fatalf("fail to check dir: %s", dir)
		return nil, err
	}
	if !existed {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			log.Fatalf("fail to mkdir: %s", dir)
			return nil, err
		}
	}

	db, err := storm.Open(path)
	if err != nil {
		log.Fatalf("fail to open dir: %s", dir)
		return nil, err
	}

	return &ItemDB{db: db}, nil
}

func (db *ItemDB) Close() error {
	return db.db.Close()
}

func (db *ItemDB) PutItem(content []byte) (item *Item, err error) {
	item = &Item{Done: false, Content: content}
	err = db.db.Save(item)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (db *ItemDB) GetItem(id int) (item *Item, err error) {
	var it Item
	err = db.db.One("ID", id, &it)
	if err != nil {
		return &it, err
	}
	return &it, nil
}

func (db *ItemDB) GetItems(n int) (item *Item, err error) {
	return item, nil
}

func (db *ItemDB) DeleteItem(id []byte) error {
	return nil
}

func (db *ItemDB) SearchItems(query []byte) ([]*Item, error) {
	return nil, nil
}
