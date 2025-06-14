package cmd

import (
	"expense-tracker-cli/internal/expense"
	"flag"
	"fmt"
)

func handleAddCommand(arguments []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	description := addCmd.String("description", "", "Description of the expense")
	amount := addCmd.Float64("amount", 0, "Amount of the expense")
	addCmd.Parse(arguments)

	if *description == "" || *amount <= 0 {
		fmt.Println("Both --description and --amount are required (amount must be provided)")
		return
	}

	expenses, err := expense.LoadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	nextID := 1
	if len(expenses) > 0 {
		nextID = expenses[len(expenses)-1].ID + 1
	}

	newExpense := expense.Expense{
		ID:          nextID,
		Description: *description,
		Amount:      *amount,
		Date:        expense.GetCurrentDate(),
	}

	expenses = append(expenses, newExpense)
	err = expense.SaveExpenses(expenses)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
		return
	}

	fmt.Printf("Expense add successfully (ID: %d)\n", newExpense.ID)
}
