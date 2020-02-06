package controllers

import (
  "encoding/json"
	"fmt"
	"github.com/kings-success/go-rest-api/pkg/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/kings-success/go-rest-api/pkg/models"
)

var NewBook models.Book;

func CreateBook(w http.ResponseWriter, r *http.Request) {
  CreateBook := &models.Book{}
  utils.ParseBody(r, CreateBook)
  b := CreateBook.CreateBook()
  res, _ := json.Marshal(b)
  w.WriteHeader(http.StatusOK)
  w.Write(res);
}

func GetBook(w http.ResponseWriter, r *http.Request){
  newBook := models.GetAllBooks()
  res, _ := json.Marshal(newBook)
  w.Header().Set("Content-Type", "pkglication/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  bookId := vars["bookId"]
  ID, err := strconv.ParseInt(bookId, 0, 0)
  if err != nil {
    fmt.Println("Error while parsing bookId")
  }

  books, _ := models.GetBookById(ID)
  res, _ := json.Marshal(books)
  w.Header().Set("Content-Type", "pkglication/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
  updateBook := &models.Book{}
  utils.ParseBody(r, updateBook)
  vars := mux.Vars(r)
  bookId := vars["bookId"]
  ID, err := strconv.ParseInt(bookId, 0, 0)
  if err != nil {
    fmt.Println("Error while parsing bookId")
  }
  book, db:= models.GetBookById(ID)
  if updateBook.Name != "" {
    book.Name = updateBook.Name
  }
  if updateBook.Author != "" {
    book.Author = updateBook.Author
  }
  if updateBook.Publication != "" {
    book.Publication = updateBook.Publication
  }
  db.Save(&book);
  res, _ := json.Marshal(book);
  w.Header().Set("Content-Type", "pkglication/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  bookId := vars["bookId"]
  ID, err := strconv.ParseInt(bookId, 0, 0)
  if err != nil {
    fmt.Println("Error while parsing bookId")
  }
  books := models.DeleteBook(ID)
  res, _ := json.Marshal(books)
  w.Header().Set("Content-Type", "pkglication/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}
