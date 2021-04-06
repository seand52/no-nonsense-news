package main

import (
	"log"
	"net/http"
	"no-nonsense-news/routes"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
  err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := httprouter.New()
	router.GET("/", routes.GetArticles)
	router.GET("/article/:articleId", routes.GetArticle)
	router.GET("/fetchNews", routes.FetchNews)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
  handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
