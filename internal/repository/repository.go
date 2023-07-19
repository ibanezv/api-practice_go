package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/rainycape/memcache"
	"gorm.io/gorm"
)

type BooksRepository interface {
	Get(ctx context.Context, id int64) (*Book, error)
}

type BooksDataBase struct {
	conn *gorm.DB
	mem  *memcache.Client
}

func NewRepository(connDb *gorm.DB, cache *memcache.Client) BooksRepository {
	return &BooksDataBase{conn: connDb, mem: cache}
}

func (db *BooksDataBase) Get(ctx context.Context, id int64) (*Book, error) {
	book := Book{}
	bookCached, err := db.mem.Get(fmt.Sprintf("book_%v", strconv.FormatInt(id, 10)))
	if err == nil {
		err := json.Unmarshal(bookCached.Value, &book)
		if err == nil {
			return &book, nil
		}
	}

	result := db.conn.Table("books").Where("id=?", id).First(&book)
	if !reflect.DeepEqual(book, Book{}) {
		payload, _ := json.MarshalIndent(&book, "", "")
		_ = db.mem.Set(&memcache.Item{
			Key:        fmt.Sprintf("book_%v", strconv.FormatInt(id, 10)),
			Value:      payload,
			Expiration: 60,
		})
	}

	return &book, result.Error
}
