package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnecet() {

	databaseUrl := "postgres://postgres:postgresql@localhost:5432/personal_web43"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Println("Koneksi database gagal", err)
		os.Exit(1)
	}

	fmt.Println("Koneksi ke database berhasil!!")
}
