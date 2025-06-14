package cmd

import (
	"expense-tracker-cli/internal/expense"
	"flag"
	"fmt"
)

func handleDeleteCommand(arguments []string) {
	deleeteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleeteCmd.Int("id", 0, "ID of the expense to delete")
	deleeteCmd.Parse(arguments)

	if *id <= 0 {
		println("--id must be provided and greater than 0")
		return
	}

	expenses, err := expense.LoadExpenses()
	if err != nil {
		println("Error loading expenses:", err.Error())
		return
	}

	found := false
	newExpenses := []expense.Expense{}
	for _, exp := range expenses {
		if exp.ID == *id {
			found = true
			continue
		}
		newExpenses = append(newExpenses, exp)
	}

	if !found {
		fmt.Printf("Expense with ID %d not found.\n", *id)
		return
	}

	err = expense.SaveExpenses(newExpenses)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
		return
	}

	fmt.Printf("Expense with ID %d deleted successfully.\n", *id)
}
