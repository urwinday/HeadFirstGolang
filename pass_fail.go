package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main3() {
	fmt.Print("Enter a grade: ")          // запрашиваем значение
	reader := bufio.NewReader(os.Stdin)   // чтение ввода с клавиатуры
	input, err := reader.ReadString('\n') // читаем вводимые данные
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)

	input = strings.TrimSpace(input)            // удаляем символ новый строки
	grade, err := strconv.ParseFloat(input, 64) // преобразование введеной строки в float64
	if err != nil {
		log.Fatal(err)
	}
	var status string // область видимости переменной в границах функции
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)

}
