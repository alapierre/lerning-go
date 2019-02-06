package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

    fmt.Println("cześć, zagraj ze mną w zgadywankę. Pomyślałem liczbę od 1 do 10, zgadnij jaka to liczba")

    var number = rand.Intn(11)
    var provided int    

    var guessed = false

	for !guessed {

		fmt.Printf("Zgadnij liczbę ")
    	fmt.Scanf("%d", &provided)   

		if provided == number {
			fmt.Println("Zgadłeś!!!")
			guessed = true
		} 
		if provided > number {
			fmt.Println("Za dużo, spróbuj jeszcze raz")
		}
		if provided < number {
			fmt.Println("Za mało, spróbuj jeszcze raz")
		}
	}
}