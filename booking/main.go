package main

import (
	"log"
	"os"

	"github.com/coroo/go-starter/routes"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	_, err := os.Stat("storage/logs")

	if os.IsNotExist(err) {
		err_0 := os.Mkdir("storage/logs", 0755)
		if err_0 != nil {
			log.Fatal(err_0)
		}
		err_1 := os.Mkdir("storage/logs/errors", 0755)
		if err_1 != nil {
			log.Fatal(err_1)
		}
		err_2 := os.Mkdir("storage/logs/informations", 0755)
		if err_2 != nil {
			log.Fatal(err_2)
		}
	}

	routes.Api()
}
