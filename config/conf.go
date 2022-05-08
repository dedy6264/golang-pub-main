package config

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func SetupConnect() (*sql.DB, error) {

	var connection = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		DBUsername, DBPassword, DBName, DBHost, DBPort, DBSSLMode)
	fmt.Println("Connection Info            : ", DBConnection, connection)

	db, err := sql.Open(DBConnection, connection)
	if err != nil {
		return db, err
	}
	fmt.Println("berhasil koneksi")

	return db, nil
}

func SetConnectionDB() {
	var err error
	db, err = SetupConnect()

	if err != nil {
		fmt.Println("Gagal Konek Database")
	}
}

func CloseConnectionDB() {
	db.Close()
}

func DbConn() *sql.DB {
	return db
}
