package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite", "demo.db") // предполагаем, что база данных называется sales.db
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть базу данных: %w", err)
	}
	defer db.Close()

	// Проверка подключения
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка проверки подключения: %w", err)
	}

	// SQL-запрос для получения данных
	query := "SELECT product, volume, date FROM sales WHERE client = ?"

	// Выполнение запроса и получение результатов
	rows, err := db.Query(query, client)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	// Обработка результатов запроса
	for rows.Next() {
		var sale Sale
		if err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}
		sales = append(sales, sale)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обработке результатов: %w", err)
	}

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
