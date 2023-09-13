package main

import(
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/gorilla/mux"

)


func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func postMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var postmovie Movie
	err := json.NewDecoder(r.Body).Decode(&postmovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postmovie.ID = strconv.Itoa(rand.Intn(100000))
	Movies = append(Movies, postmovie)
	json.NewEncoder(w).Encode(postmovie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range Movies {
		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			json.NewEncoder(w).Encode(Movies)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range Movies {
		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			var Upmovie Movie
			err := json.NewDecoder(r.Body).Decode(&Upmovie)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			Upmovie.ID = params["id"]
			Movies = append(Movies, Upmovie)
			json.NewEncoder(w).Encode(Upmovie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}
