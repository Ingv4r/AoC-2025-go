package main

import (
	"AoC-2025-go/internal/src"
	"AoC-2025-go/internal/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envPath := ".env"

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("не удалось загрузить .env файл: %v", err)
	}
	sessionCookie := os.Getenv("AOC_SESSION")

	input, err := utils.FetchInput(2025, 1, sessionCookie)
	if err != nil {
		fmt.Printf("ошибка: %v\n", err)
		return
	}

	zeros, err := src.GetZeros(input)
	if err != nil {
		log.Panicf("ошибка: %v\n", err)
	}
	println(zeros)
}
