package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"no-nonsense-news/routes"
)

func main() {
	router := httprouter.New()
	router.GET("/", routes.GetArticles)
	router.GET("/article/:articleId", routes.GetArticle)
	router.GET("/fetchNews", routes.FetchNews)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
