package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func main() {
// 	PATH := "localhost:8080"
// 	router := gin.Default()

// 	router.GET("/", SayHello)
// 	router.GET("/health", GetHealth)

// 	router.Run(PATH)
// }

// func SayHello(context *gin.Context) {
// 	context.IndentedJSON(
// 		http.StatusOK,
// 		"Hello. I'm Up on port 8080",
// 	)
// }

// func GetHealth(context *gin.Context) {
// 	context.IndentedJSON(
// 		http.StatusOK,
// 		"OK, I'M UP",
// 	)
// }
