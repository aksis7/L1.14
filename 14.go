package main

import (
	"fmt"
)

func determineType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan any:
		fmt.Println("Тип: channel")
	default:
		fmt.Println("Неизвестный тип:", v)
	}
}

func main() {
	var x interface{}

	// Пример с типом int
	x = 42
	determineType(x)

	// Пример с типом string
	x = "Hello, Go!"
	determineType(x)

	// Пример с типом bool
	x = true
	determineType(x)

	// Пример с типом channel
	x = make(chan interface{})
	determineType(x)

	// Пример с неизвестным типом
	x = 3.14
	determineType(x)
}
