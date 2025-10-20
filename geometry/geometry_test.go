package geometry

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Функция для округления числа до сотых и преобразования в строку
func roundToTwoDecimalPlaces(h float64) string {
	return fmt.Sprintf("%.2f", h)
}

func TestHypotenuse(t *testing.T) {
	items := []struct {
		a, b     float64
		expected string
	}{
		{0, 0, "0.00"},
		{0, 9, "9.00"},
		{3, 4, "5.00"},
		{10, 21, "23.26"},
		{56, 37, "67.12"},
		{102, 67, "122.04"},
		{34, 17, "38.01"},
		{3456, 1089, "3623.51"},
		{478, 201, "518.54"},
	}

	for _, item := range items {
		h := Hypotenuse(item.a, item.b)
		assert.Equal(t, item.expected, roundToTwoDecimalPlaces(h))
	}
}

func TestHypotenuseSimple(t *testing.T) {
	h := Hypotenuse(3, 4)

	assert.Equal(t, 5.0, h)
}
