package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAddCommand(os.Args[2:])
	case "list":
		handleListCommand()
	case "delete":
		handleDeleteCommand(os.Args[2:])
	case "summary":
		handleSummaryCommand(os.Args[2:])
	case "help":
		showUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showUsage()
	}
}

func showUsage() {
	fmt.Println("Expense Tracker CLI")
	fmt.Println("Usage:")
	fmt.Println("   add --description <description> --amount <amount>")
	fmt.Println("   list")
	fmt.Println("   delete id <id>")
	fmt.Println("   summary [--month <month>]")
	fmt.Println("   help")
}
