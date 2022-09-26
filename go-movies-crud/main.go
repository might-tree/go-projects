//Using structs and slices, not using database

package main

import (
	"encoding/json" //to encode data into json when we send it to Postman
	"fmt"           //for output
	"log"           //to log out any errors
	"math/rand"     //to generate random ID for newly added elements"
	"net/http"      //to create a server in golang
	"strconv"       //to convert ID from integer to string

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json: "isbn:`
	Title    string    `json: "title"`
	Director *Director `json: "director"` //Pointer
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) { //params: response sent from this func, pointer of the request that we will send from Postman to func
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) //Encoding entire slice that we have into json (when we add a new element, we are simply appending to an existing slice)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) { //params: response sent from this func, pointer of the request that we will send from Postman to func
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) { ////params: response sent from this func, pointer of the request that we will send from Postman to func
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
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting content type
	params := mux.Vars(r)                              //params

	//Loop over elements, delete the movie with the id sent, add new movie sent to body of Postman
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", ISBN: "435678", Title: "Dunkirk", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}}) //to test if server is working

	r.HandleFunc("/movies", getMovies).Methods("GET") //function definitions and routes
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r)) //localhost:8000
}
