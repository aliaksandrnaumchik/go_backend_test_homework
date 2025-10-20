package main

import (
	"fmt"
	"math/rand"
)

func Winner(rnd *rand.Rand, bets map[int]string) string {
	// Генерируем случайное число от 0 до 99
	randomNumber := rnd.Int() % 100

	// Инициализируем переменные для отслеживания победителя
	minDifference := 100 // Максимальное возможное значение разницы
	winner := ""         // Имя победителя

	// Проходим по всем ставкам игроков
	for number, player := range bets {
		// Вычисляем разницу
		difference := number - randomNumber

		// Если разница отрицательная, меняем знак
		if difference < 0 {
			difference = -difference
		}

		// Если нашли меньшую разницу - обновляем победителя
		if difference < minDifference {
			minDifference = difference
			winner = player
		}
	}

	return winner
}

func main() {
	rnd := rand.New(rand.NewSource(1001))
	// кто какое число загадал
	bets := map[int]string{
		20: "Маша",
		34: "Игорь",
		77: "Олег",
		51: "Света",
		2:  "Саша",
	}

	fmt.Println("Победитель:", Winner(rnd, bets))
}
