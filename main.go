package main

import (
	"api_gateway/usecase"
	"fmt"
)

func main() {

	loginService := usecase.NewLogin()

	username := "admin"
	password := "admin123"
	isAuthenticated := loginService.Authentifikasi(username, password)

	// Menampilkan hasil autentikasi
	fmt.Println("Authentifikasi:", isAuthenticated)

}
