package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

// Expense struct definition
type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}

const dataFile = "expenses.json"

func main() {
	args := os.Args

	if len(args) < 2 {
		showUsage()
		return
	}

	command := args[1]

	switch command {
	case "add":
		handleAdd(args[2:])
	case "list":
		handleList()
	case "delete":
		fmt.Println("Delete command triggered")
	case "summary":
		fmt.Println("Summary command triggered")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showUsage()
	}
}

func handleAdd(arguments []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	description := addCmd.String("description", "", "Expense description")
	amount := addCmd.Float64("amount", 0.0, "Expense amount")

	addCmd.Parse(arguments)

	if *description == "" || *amount <= 0 {
		fmt.Println("Both --description and --amount are required (amount must be provided)")
		return
	}

	expenses, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	nextID := 1
	if len(expenses) > 0 {
		nextID = expenses[len(expenses)-1].ID + 1
	}

	newExpense := Expense{
		ID:          nextID,
		Description: *description,
		Amount:      *amount,
		Date:        getCurrentDate(),
	}

	expenses = append(expenses, newExpense)
	err = saveExpenses(expenses)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
		return
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", newExpense.ID)
}

func loadExpenses() ([]Expense, error) {
	var expenses []Expense

	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return expenses, nil
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return expenses, nil
	}

	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func saveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}

func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func handleList() {
	expenses, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Printf("%-4s %-12s %-20s %s\n", "ID", "Date", "Description", "Amount")
	for _, exp := range expenses {
		fmt.Printf("%-4d %-12s %-20s $%.2f\n", exp.ID, exp.Date, exp.Description, exp.Amount)
	}
}

func showUsage() {
	fmt.Println("Usage:")
	fmt.Println("   add --description <description> --amount <amount>")
	fmt.Println("   list")
	fmt.Println("   delete")
	fmt.Println("   summary")
}
