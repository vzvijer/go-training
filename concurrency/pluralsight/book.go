package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

func (b Book) String() string {
	return fmt.Sprintf("Title:\t\t%q\nAuthor:\t\t%q\nYear:\t\t%v\n", b.Title, b.Author, b.Year)
}

var books = []Book{
	{
		ID:     1,
		Title:  "Things Fall Apart",
		Author: "Chinua Achebe",
		Year:   1958,
	},
	{
		ID:     2,
		Title:  "Fairy tales",
		Author: "Hans Christian Andersen",
		Year:   1836,
	},
	{
		ID:     3,
		Title:  "The Divine Comedy",
		Author: "Dante Alighieri",
		Year:   1315,
	},
	{
		ID:     4,
		Title:  "The Epic Of Gilgamesh",
		Author: "Unknown",
		Year:   -1700,
	},
	{
		ID:     5,
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Year:   1813,
	},
	{
		ID:     6,
		Title:  "Molloy, Malone Dies, The Unnamable, the trilogy",
		Author: "Samuel Beckett",
		Year:   1952,
	},
	{
		ID:     7,
		Title:  "The Decameron",
		Author: "Giovanni Boccaccio",
		Year:   1351,
	},
	{
		ID:     8,
		Title:  "The Stranger",
		Author: "Albert Camus",
		Year:   1942,
	},
	{
		ID:     9,
		Title:  "Ficciones",
		Author: "Jorge Luis Borges",
		Year:   1965,
	},
	{
		ID:     10,
		Title:  "Don Quijote De La Mancha",
		Author: "Miguel de Cervantes",
		Year:   1610,
	},
}
