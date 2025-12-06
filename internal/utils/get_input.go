package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FetchInput(year, day int, sessionCookie string) ([]string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {return nil, fmt.Errorf("ошибка создания запроса: %v", err)}

	request.AddCookie(&http.Cookie{
		Name: "session",
		Value: sessionCookie,
	})
	
	response, err := client.Do(request)
	if err != nil {return nil, fmt.Errorf("ошибка отправки запроса: %v", err)}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("неверный статус ответа: %d %s", response.StatusCode, response.Status)
    }

	body, err := io.ReadAll(response.Body)
    if err != nil {return nil, fmt.Errorf("ошибка чтения ответа: %v", err)}

	str := string(body)
	fmt.Printf("Получено %d символов\n", len(str))
	result := strings.Split(strings.TrimSpace(str), "\n")
    
    return result, nil
}