package main

import (
	"fmt"

	"gopkg.in/yaml.v3" // импортируем пакет для работы с YAML
)

// Artist содержит данные об артисте
type Artist struct {
	ID    int      `yaml:"id"`
	Name  string   `yaml:"name"`
	Genre string   `yaml:"genre"`
	Songs []string `yaml:"songs"`
}

func main() {
	// создаём переменную типа Artist, чтобы конвертировать эти данные в YAML
	artist := Artist{
		ID:    1,
		Name:  "30 seconds to Mars",
		Genre: "rock",
		Songs: []string{
			`The Kill`,
			`A Beautiful Lie`,
			`Attack`,
			`Live Like A Dream`,
		},
	}

	// сериализуем данные из переменной artist в слайс байт
	yamlData, err := yaml.Marshal(&artist)
	if err != nil {
		fmt.Println("ошибка при сериализации в yaml:", err)
		return
	}

	var newArtist Artist
	err = yaml.Unmarshal(yamlData, &newArtist)
	if err != nil {
		fmt.Println("ошибка при десериализации:", err)
		return
	}

	fmt.Println("ID:", newArtist.ID)
	fmt.Println("Name:", newArtist.Name)
	fmt.Println("Genre:", newArtist.Genre)
	fmt.Printf("Songs: %v\n", newArtist.Songs)
}
