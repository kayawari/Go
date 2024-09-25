package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux" // パスマッチングを行うためのライブラリ"
	"log"
	"math/rand"
	"net/http"
	"strconv" // 整数を文字列に変換する
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[index], movies[index+1:]...) // POINT!!!!
			break
		}
	}

	json.NewEncoder(w).Encode(movies) // レスポンスを返す時に必要なもの。encodeはここで使う
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// ...............
	// ...............
	// ...............
	// ...............
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "111-1111", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "222-2222", Title: "Movie Two", Director: &Director{Firstname: "Paul", Lastname: "George"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
