package main

import "fmt"

// Hero описывает главного героя игры
type Hero struct {
	Name    string // имя героя
	Health  int    // здоровье
	Damage  int    // урон
	Defense int    // защита
}

// ChangeHealth изменяет здоровье героя на указанное значение.
// Для структуры Hero добавьте метод ChangeHealth(diff int).
// Он должен понижать уровень здоровья героя на указанную величину при отрицательном значении параметра diff
// и повышать при положительном.
// вставьте метод ChangeHealth

func (hero *Hero) ChangeHealth(diff int) {
	hero.Health += diff
}

func main() {
	hero := &Hero{
		Name:   "Илья Муромец",
		Health: 100500,
	}
	for i := -100; i <= 150; i += 20 {
		hero.ChangeHealth(i)
	}
	fmt.Println(hero.Health) // должна вывести 100760
}
