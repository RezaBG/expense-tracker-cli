package cmd

import (
	"expense-tracker-cli/internal/expense"
	"fmt"
)

func handleListCommand() {
	expenses, err := expense.LoadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	if len(expenses) == 1 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Printf("%-4s %-12s %-20s %s\n", "ID", "Date", "Description", "Amount")
	for _, exp := range expenses {
		fmt.Printf("%-4d %-12s %-20s %.2f\n", exp.ID, exp.Date, exp.Description, exp.Amount)
	}
}
