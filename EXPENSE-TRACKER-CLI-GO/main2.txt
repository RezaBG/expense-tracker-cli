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
		handleDelete(args[2:])
	case "summary":
		handleSummary(args[2:])
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

func handleDelete(arguments []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "Expense ID to delete")

	deleteCmd.Parse(arguments)

	if *id <= 0 {
		fmt.Println("Expense ID must be greater than 0")
		return
	}

	expenses, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	// Find and delete the expense
	found := false
	newExpenses := []Expense{}

	for _, exp := range expenses {
		if exp.ID == *id {
			found = true
			continue // Skip the expense to delete it
		}
		newExpenses = append(newExpenses, exp)
	}

	if !found {
		fmt.Printf("Expense with ID %d not found.\n", *id)
		return
	}

	err = saveExpenses(newExpenses)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
		return
	}

	fmt.Printf("Expense with ID %d deleted successfully.\n", *id)
}

func handleSummary(arguments []string) {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := summaryCmd.Int("month", 0, "Month number (1-12)")

	summaryCmd.Parse(arguments)

	expenses, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found for summary.")
		return
	}

	total := 0.0

	for _, exp := range expenses {
		if *month > 0 {
			expensesTime, err := time.Parse("2006-01-02", exp.Date)
			if err != nil {
				continue
			}
			if int(expensesTime.Month()) != *month {
				continue
			}
		}
		total += exp.Amount
	}

	// Print results
	if *month > 0 {
		fmt.Printf("Total expenses for month %d: $%.2f\n", *month, total)
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}
}

func showUsage() {
	fmt.Println("Usage:")
	fmt.Println("   add --description <description> --amount <amount>")
	fmt.Println("   list")
	fmt.Println("   delete")
	fmt.Println("   summary")
}
