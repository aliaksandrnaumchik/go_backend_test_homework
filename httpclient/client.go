package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
		return
	}
	// читаем тело ответа
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	fmt.Println(string(body))
}
