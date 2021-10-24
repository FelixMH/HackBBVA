package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env") // carga el archivo .env que contiene las credenciales
var ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", //define la conexion a la bd
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))

func GetDB() (*sql.DB, error) { //conecta mysql con go
	return sql.Open("mysql", ConnectionString)
}