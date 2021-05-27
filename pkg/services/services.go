package services

import "fmt"

func GetUserBalance(userId int) {
	fmt.Println(userId)
}

func ChangeBalance(userId int, amount float64, direction string) {
	fmt.Println(direction)
}

func Transaction(withdrawUserId int, refillUserId int, amount float64) {
	fmt.Println(amount)
}
