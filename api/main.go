package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Construir cadena de conexión
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)

	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal("No se pudo conectar a la BD:", err)
	}
	defer db.Close()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Funcionando"))
	})

	fmt.Println("API corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
