package main

import (
	"net/http"

	"github.com/go-martini/martini"
)

type Book struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

func Something(req *http.Request, res http.ResponseWriter, params martini.Params) (int, string) {
	return http.StatusOK, params["param"]
}

func main() {
	m := martini.New()
	router := martini.NewRouter()

	router.NotFound(func() (int, []byte) {
		return http.StatusNotFound, nil
	})

	m.Use(martini.Logger())

	router.Get("/something/:param", Something)

	m.MapTo(router, (*martini.Routes)(nil))
	m.Action(router.Handle)
	m.Run()
}
