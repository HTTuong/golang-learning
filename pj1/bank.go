package main

import (
	"errors"
	"first-app/fileops"
	"fmt"
	"os"
)

const inflationRate = 2.5
const accountBalanceFile = "balance.txt"

func main() {

	var choice int
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)


	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic("Cannot continue, sorry")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()
		fmt.Print("Enter your option: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice: ", choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is", accountBalance)
		case 2: 
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if (depositAmount <= 0) {
				fmt.Print("The amount is invalid. Please try again!")
				continue
			}

			accountBalance +=  depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance  )
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)
	
			if (withdrawAmount <= 0) {
				fmt.Print("Invalid amount. Please try again!")
				continue
			}
	
			if (withdrawAmount > accountBalance) {
				fmt.Print("You cannot withdrawn the amount that you are currently have!")
				continue
			}
	
			accountBalance -=  withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance  )
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing this bank")
			return 
		}
	}



	// revenue, errRevenue := getInput("Revenue: ") 
	// expenses, errExpenses := getInput("Expenses: ") 
	// taxRate, errTaxRate := getInput("Tax Rate: ") 

	// if errRevenue != nil || errExpenses != nil || errTaxRate != nil {
	// 	fmt.Println("Invalid entered value")
	// 	fmt.Println("The program will shut down!")
	// 	return
	// }
	

	// ebt , profit , ratio := calculateFinancials(revenue, expenses, taxRate)

	// fmt.Printf("EBT: %.1f\n", ebt)
	// fmt.Printf("Profit: %.1f\n", profit)
	// fmt.Printf("Ratio: %.1f\n", ratio)

	// saveResultInFile(ebt, profit, ratio)

}



func saveResultInFile(ebt float64, profit float64, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
} 

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return revenue, expenses, ratio
}

func getInput(message string) (inputValue float64, err error) {
	fmt.Print(message)
	fmt.Scan(&inputValue)

	if inputValue < 0 {
		return inputValue, errors.New("negative input. please try again")
	}

	if inputValue == 0 {
		return 0, errors.New("0 is not allowed. please try again")
	}
	
	return inputValue, nil
}