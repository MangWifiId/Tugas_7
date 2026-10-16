
package main

import (
	"fmt"
	"log"

	"Pelatihan_7/database"
	"Pelatihan_7/router"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}
	defer db.Close()

	fmt.Println("Terhubung ke database PostgreSQL")

	r := router.NewRouter(db)

	fmt.Println("Server Gin berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
