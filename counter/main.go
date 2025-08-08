package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Counter возвращает строку для записи в файл.
func Counter(count int, t time.Time) string {
	// Форматируем время в читаемый формат
	formattedTime := t.Format("02.01.2006")

	// Создаем строку с результатом
	return fmt.Sprintf("%d %s", count, formattedTime)
}

// Limits возвращает количество дней и запусков.
func Limits() (int, int, error) {
	// получаем имя программы
	app, err := os.Executable()
	if err != nil {
		return 0, 0, err
	}
	// получаем путь и имя текстового файла
	name := filepath.Join(filepath.Dir(app), "data.txt")
	if _, err = os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			// записываем начальные значения
			out := Counter(1, time.Now())
			err = os.WriteFile(name, []byte(out), 0644)
			return 0, 1, err
		}
		return 0, 0, err
	}
	var data []byte
	// читаем файл
	data, err = os.ReadFile(name)
	if err != nil {
		return 0, 0, err
	}
	counter, t, err := ParseCounter(string(data))
	if err != nil {
		return 0, 0, err
	}
	// сохраняем в файл новое значение счётчика
	counter++
	// время записывается без измененений
	if err = os.WriteFile(name, []byte(Counter(counter, t)), 0644); err != nil {
		return 0, 0, err
	}
	duration := time.Since(t)
	// считаем количество дней
	return int(duration.Hours()) / 24, counter, nil
}

// ParseCounter разбирает информацию из файла.
// Функция возвращает значение счётчика и дату первого запуска.
func ParseCounter(input string) (int, time.Time, error) {
	// Разделяем входную строку на части
	parts := strings.Split(input, " ")

	if len(parts) != 2 {
		return 0, time.Time{}, fmt.Errorf("неверный формат строки: ожидается 'число дата'")
	}

	// Парсим число
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("ошибка при парсинге числа: %w", err)
	}

	// Парсим дату
	t, err := time.Parse("02.01.2006", parts[1])
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("ошибка при парсинге даты: %w", err)
	}

	return count, t, nil
}

func main() {
	days, counter, err := Limits()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Количество дней: %d\nКоличество запусков: %d\n", days, counter)
	// устанавливаем лимит в 14 дней или 50 запусков
	if days > 14 || counter > 50 {
		fmt.Println("Запросите новую версию")
		return
	}
	fmt.Println("Программа готова к работе")
}
