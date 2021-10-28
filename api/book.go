package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}

var books = map[string]Book{
	"978-1742611778": {Title: "Area 7", Author: "Matthew Reilly", ISBN: "978-1742611778"},
	"978-0066214122": {Title: "Pray", Author: "Michael Crichton", ISBN: "978-0066214122"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}
}

func AllBooks() []Book {
	return nil
}

func GetBook(isbn string) (Book, bool) {
	return Book{}, false
}

func CreateBook(book Book) (string, bool) {
	return "", false
}

func UpdateBook(isbn string, book Book) bool {
	return false
}

func DeleteBook(isbn string) {

}

func (b Book) ToJSON() []byte {
	mj, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return mj
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(b)
}

func FromJSON(data []byte) Book {
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

//func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
//	b, err := json.Marshal(Books)
//	if err != nil {
//		panic(err)
//	}
//
//	w.Header().Add("Content-Type", "application/json; charset=utf-8")
//	w.Write(b)
//}
