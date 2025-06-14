package expense

import (
	"encoding/json"
	"os"
	"time"
)

const dataFile = "expenses.json"

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}

func LoadExpenses() ([]Expense, error) {
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

// SaveExpenses saves the expenses to JSON file
func SaveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}

// GetCurrentDate returns today's date
func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}
