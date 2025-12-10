package main

import (
	"AoC-2025-go/internal/src"
	"fmt"
	"AoC-2025-go/internal/utils"
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

	input, err := utils.FetchInput(2025, 3, sessionCookie)
	if err != nil {
		fmt.Printf("ошибка: %v\n", err)
		return
	}

	fmt.Printf("Sum of voltages: %d", src.MaxVoltage(input))
}
