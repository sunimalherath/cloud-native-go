package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) ToJSON() []byte {
	mj, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return mj
}

func (b Book) FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	{Title: "Area 7", Author: "Matthew Reilly", ISBN: "978-1742611778"},
	{Title: "Pray", Author: "Michael Crichton", ISBN: "978-0066214122"},
}

func BooksHandlFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
