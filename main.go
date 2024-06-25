package main

import (
	"go-web-native/config"
	categorycontrollers "go-web-native/controllers/category_controllers"
	homecontrollers "go-web-native/controllers/home_controllers"
	productcontrollers "go-web-native/controllers/product_controllers"
	"log"
	"net/http"
)

func main() {
	config.Connection()

	// 1. Homepage
	http.HandleFunc("/", homecontrollers.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categorycontrollers.Index)
	http.HandleFunc("/categories/add", categorycontrollers.Add)
	http.HandleFunc("/categories/edit", categorycontrollers.Edit)
	http.HandleFunc("/categories/delete", categorycontrollers.Delete)

	// 3. Products
	http.HandleFunc("/products", productcontrollers.Index)
	http.HandleFunc("/products/add", productcontrollers.Add)
	http.HandleFunc("/products/detail", productcontrollers.Detail)
	http.HandleFunc("/products/edit", productcontrollers.Edit)
	http.HandleFunc("/products/delete", productcontrollers.Delete)

	log.Println("Server running on port : 8080")
	http.ListenAndServe(":8080", nil)
}
