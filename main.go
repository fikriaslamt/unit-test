package main

import (
	"sesi_8/app"
	"sesi_8/config"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	app.StartApplication()
}
