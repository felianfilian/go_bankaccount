package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const saveFile = "balance.txt"

func main() {
	var option int
	var balance float64 = readBalance()
	var amount float64

	for {
		showOptions(balance)
		fmt.Scan(&option)

		// switch option {
		// 	case 4:
		// 		return
		// }

		if option == 1 {

		} else if option == 2 {
			fmt.Print("How much: ")
			fmt.Scan(&amount)
			balance += amount
			writeToFile(balance)

		} else if option == 3 {
			fmt.Print("How much: ")
			fmt.Scan(&amount)
			if amount > balance {
				fmt.Printf("\nNot enough cash\n\n")
			} else {
				balance -= amount
			}
			writeToFile(balance)
		} else if option == 4 {
			fmt.Println("Goodbye")
			break
		}
	}

}

func showOptions(balance float64) {
	fmt.Print("Actual Balace: ")
	fmt.Print(balance)
	fmt.Println("\nMain Menu")
	fmt.Println("1) check balance")
	fmt.Println("2) deposit money")
	fmt.Println("3) withdraw money")
	fmt.Println("4) exit")
	fmt.Print("choice: ")
}

func writeToFile(value float64) {
	txtBalance := fmt.Sprint(value)
	os.WriteFile(saveFile, []byte(txtBalance), 0644)
}

func readBalance() (float64, error) {
	data, err := os.ReadFile(saveFile)
	if(err != nil) {
		return 1000, errors.New("failed to find file")
	}
	txtBalance := string(data)
	balance, err := strconv.ParseFloat(txtBalance, 64)
	if(err != nil) {
		return 1000, errors.New("failed to parse balance")
	}

	return balance, nil
}
