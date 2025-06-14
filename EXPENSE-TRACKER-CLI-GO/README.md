# Expense Tracker CLI (Golang V2 Professional Version)

A simple but fully production-ready expense tracker CLI application written in Go.

This project demonstrates:

- Clean Go project architecture
- Separation of concerns (business logic vs CLI interface)
- File-based JSON storage
- Full CRUD operations for expense tracking

---

## Features

Add new expenses  
List expenses  
Delete expenses by ID  
Summarize total expenses  
Summarize monthly expenses  
Fully extensible architecture

---

## Usage

```bash
# Add new expense
go run . add --description "Lunch" --amount 20

# List all expenses
go run . list

# Delete expense by ID
go run . delete --id 2

# Show total expenses
go run . summary

# Show monthly expenses
go run . summary --month 6
```

## Project Submission

This project was built as part of the [Expense Tracker CLI project challenge on roadmap.sh](https://roadmap.sh/projects/expense-tracker-cli).
