package main

import (
    "encoding/json"
    "net/http"
	"path"
	"fmt"
)

func find(x string) int {
    for i, book := range books {
        if x == book.Id {
            return i
        }
    }
    return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	if id == "book"{
		dataJson, err := json.Marshal(books)
		checkError("Parse error", err)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
	}
    checkError("Parse error", err)
    i := find(id)
    if i == -1 {
        return
    }
    dataJson, err := json.Marshal(books[i])
    w.Header().Set("Content-Type", "application/json")
    w.Write(dataJson)
    return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    book := Book{}
    json.Unmarshal(body, &book)
    books = append(books, book)
    w.WriteHeader(200)
    return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	i := find(id)
	bookRef := &books[i]

	len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    book := Book{}
	json.Unmarshal(body, &book)
	
	bookRef.Title = isValidUpdate(bookRef.Title, book.Title)
	bookRef.Edition = isValidUpdate(bookRef.Edition, book.Edition)
	bookRef.Copyright = isValidUpdate(bookRef.Copyright, book.Copyright)
	bookRef.Language = isValidUpdate(bookRef.Language, book.Language)
	bookRef.Pages = isValidUpdate(bookRef.Pages, book.Pages)
	bookRef.Author = isValidUpdate(bookRef.Author, book.Author)
	bookRef.Publisher = isValidUpdate(bookRef.Publisher, book.Publisher)

    w.WriteHeader(200)
    return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	i := find(id)
	books = append(books[:i], books[i+1:]...)
    w.WriteHeader(200)
    return
}

func isValidUpdate(current string, new string) string {
	if new != "" {
		return new
	}

	return current
}