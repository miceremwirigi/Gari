package routes

import (
	"Gari/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/cars/", handlers.GetCars)
	http.HandleFunc("/add/car/", handlers.AddCar)
	http.HandleFunc("/delete/car/", handlers.DeleteCar)
	http.HandleFunc("/update/car/", handlers.UpdateCar)
}
