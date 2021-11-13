// guess - игра, в которой игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // получаем текущую дату и время в формате целого числа
	fmt.Println(seconds)
	rand.Seed(seconds) // функция генератор случайных чисел
	target := rand.Intn(100) + 1
	//fmt.Println(target)
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin) // для чтения ввода с клавиатуры

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n') // Прочитать данные введеные пользователем до нажатия Enter
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  //удаление символа
		guess, err := strconv.Atoi(input) // входная строка преабразуется в целое число
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true // число угадоно сообщение о пройгрыше выводить не нужно
			fmt.Println("Good job! You guessed it!")
			break
		}

		if !success {
			fmt.Println("Sorry, you didn't guess my number. It was:", target)
		}
	}

}
