package routes

import (
	"github.com/king-success/go-rest-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func (router *mux.Router) {
  router.HandleFunc("/book/", controllers.CreateBook).Method("POST");
  router.HandleFunc("/book/", controllers.GetBook).Method("GET");
  router.HandleFunc("/book/{bookId}", controllers.GetBookById).Method("GET");
  router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Method("PUT");
  router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Method("DELETE");
}
