package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConnectionDB is the string connection to the database
	StringConnectionDB = ""

	// Port is the port where the application will run
	Port = 0
)

// ToLoad loads the application configuration
func ToLoad() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("DB_PORT"))
	if erro != nil {
		Port = 9000
	}

	StringConnectionDB = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),     // Usuário do banco
		os.Getenv("DB_PASSWORD"), // Senha do banco
		os.Getenv("DB_HOST"),     // Host do banco (localhost ou IP do container)
		os.Getenv("DB_PORT"),     // Porta do banco (3306)
		os.Getenv("DB_NAME"),     // Nome do banco
	)
}
