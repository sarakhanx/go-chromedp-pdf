package dotenv

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Simple := os.Getenv("SIMPLE")
	if err != nil {
		fmt.Println("something went wrong", err)
	}
	fmt.Println("SIMPLE Text = ", Simple)
}
