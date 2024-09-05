package main

import "fmt"

func main() {
	var option int
	for {
		showOptions()
		fmt.Scan(&option)

		// switch option {
		// 	case 4:
		// 		return
		// }

		if(option == 1) {

		} else if(option == 2) {
			
		} else if(option == 4) {
			fmt.Println("Goodbye")
			break
		}
	}
	
}

func showOptions() {
	fmt.Println("\nMain Menu")
	fmt.Println("1) check balance")
	fmt.Println("2) deposit money")
	fmt.Println("3) withdraw money")
	fmt.Println("4) exit")
	fmt.Print("choice: ")
}