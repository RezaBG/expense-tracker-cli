package cmd

import (
	"expense-tracker-cli/internal/expense"
	"flag"
	"fmt"
	"time"
)

func handleSummaryCommand(arguments []string) {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := summaryCmd.Int("month", 0, "Month number (1-12) to filter expenses")
	summaryCmd.Parse(arguments)

	expenses, err := expense.LoadExpenses()
	if err != nil {
		println("Error loading expenses:", err.Error())
		return
	}

	if len(expenses) == 0 {
		println("No expenses found.")
		return
	}

	total := 0.0
	for _, exp := range expenses {
		if *month > 0 {
			expenseTime, err := time.Parse("2006-01-02", exp.Date)
			if err != nil {
				continue
			}
			if int(expenseTime.Month()) != *month {
				continue
			}
		}
		total += exp.Amount
	}

	if *month > 0 {
		fmt.Printf("Total expenses for month %d: %.2f\n", *month, total)
	} else {
		fmt.Printf("Total expenses: %.2f\n", total)
	}
}
