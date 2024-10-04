package main

import (
	"fmt"
	Project1 "goprojects/Functions"
)

func main() {
	var number int
	var selection int

	fmt.Println("Welcome to Animal Farm Program 1.0.")

	for {
		fmt.Println("Please choose an option:")
		fmt.Println("1) Add new information")
		fmt.Println("2) Take info for selected category")
		fmt.Println("3) Update an information")
		fmt.Println("4) Take info for all farm")
		fmt.Println("5) Exit")

		_, err := fmt.Scanln(&number)
		if err != nil || number < 1 || number > 5 {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			continue
		}

		switch number {
		case 1:
			Project1.AddNewInformation()

		case 2:
			fmt.Println("Please select a category:")
			fmt.Println("1) Cattle")
			fmt.Println("2) Chickens")
			fmt.Println("3) Vets")
			fmt.Println("4) Feeds")
			fmt.Println("5) Butchers")
			fmt.Println("6) Feed Sellers")
			fmt.Println("7) Monthly Income")
			fmt.Println("8) Monthly Expense")

			_, err := fmt.Scanln(&selection)
			if err != nil || selection < 1 || selection > 8 {
				fmt.Println("Invalid input. Please select a number between 1 and 8.")
				continue
			}

			switch selection {
			case 1:
				fmt.Println("Cattle Info:")
				Project1.PrintCattleInfo()

			case 2:
				fmt.Println("Chicken Info:")
				Project1.PrintChickenInfo()

			case 3:
				fmt.Println("Vet Info:")
				Project1.PrintVetInfo()

			case 4:
				fmt.Println("Feed Info:")
				Project1.PrintFeedInfo()

			case 5:
				fmt.Println("Butcher Info:")
				Project1.PrintButcherInfo()

			case 6:
				fmt.Println("Feed Seller Info:")
				Project1.PrintFeedSellerInfo()

			case 7:
				fmt.Println("Monthly Income Info:")
				Project1.PrintMonthlyIncomeInfo()

			case 8:
				fmt.Println("Monthly Expense Info:")
				Project1.PrintMonthlyExpenseInfo()
			}

		case 3:
			Project1.UpdateInformation()

		case 4:
			Project1.TakeInfoFarm()

		case 5:
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid option. Please enter a number between 1 and 5.")
		}
	}
}
