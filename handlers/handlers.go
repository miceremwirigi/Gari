package handlers

import (
	"Gari/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Add("Error", "Wrong Method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Success", "Right Method")

	io.WriteString(w, "Welcome")
	w.WriteHeader(http.StatusOK)
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Add("Error", "Wrong Method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Success", "Right Method")

	json.NewEncoder(w).Encode(models.Cars)

	// io.WriteString(w, "Welcome")
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Add("Error", "Wrong Method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Success", "Right Method")

	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)

	models.Cars = append(models.Cars, car)

	// io.WriteString(w, "Welcome")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Addeded car sussessfully"))

}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Add("Error", "Wrong Method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()

	id, _ := strconv.Atoi(query.Get("id"))

	for index, car := range models.Cars {
		if car.ID == id {
			models.Cars = append(models.Cars[:index], models.Cars[index+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted car sussessfully"))
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		w.Header().Add("Error", "Wrong Method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updatedCar models.Car
	json.NewDecoder(r.Body).Decode(&updatedCar)

	query := r.URL.Query()

	id, _ := strconv.Atoi(query.Get("id"))

	for index, car := range models.Cars {
		if car.ID == id {
			models.Cars[index] = updatedCar
		}
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated car sussessfully"))

}
