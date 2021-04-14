package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		go func(id int) {
			if book, ok := queryCache(id); ok {
				fmt.Println("Cache:")
				fmt.Println(book)
			}
		}(id)

		go func(id int) {
			if book, ok := queryDatabase(id); ok {
				fmt.Println("Database:")
				fmt.Println(book)
			}
		}(id)

		//time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (Book, bool) {
	book, ok := cache[id]
	return book, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, book := range books {
		if book.ID == id {
			cache[id] = book
			return book, true
		}
	}

	return Book{}, false
}
