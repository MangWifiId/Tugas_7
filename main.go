package main

import (
	"fmt"
	"log"

	"Pelatihan_7/database"
	"Pelatihan_7/router"
	"Pelatihan_7/todo"
)

func main() {
	// Koneksi ke database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan koneksi SQL: ", err)
	}
	defer sqlDB.Close()

	fmt.Println("Terhubung ke database PostgreSQL")

	// Migrasi tabel Todo
	if err := db.AutoMigrate(&todo.Todo{}); err != nil {
		log.Fatal("Gagal migrasi database: ", err)
	}

	fmt.Println("Migrasi database berhasil")

	// Setup router Gin
	r := router.NewRouter(db)

	// Menjalankan server
	fmt.Println("Server Gin berjalan di http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
