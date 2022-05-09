package main

import (
	"atmcase/app"
	"atmcase/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("gagal baca file .env")
	}

	dataAccount := model.ReadDataAccountJson(os.Getenv("FILE_JSON"))

	app.StartApp(dataAccount)
}
