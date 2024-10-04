package Project1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Cattle4 struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Birthdate   string    `json:"birthdate"`
	Deathdate   string    `json:"deathdate"`
	EarNumber   int       `json:"ear_number"`
	Gender      string    `json:"gender"`
	MotherID    int       `json:"motherId"`
	FatherID    int       `json:"fatherId"`
	Race        string    `json:"race"`
	Color       string    `json:"color"`
	Height      float64   `json:"height"`
	ChestSize   []float64 `json:"chest_size"`
	BirthNumber int       `json:"birth_number"`
}

type Chicken4 struct {
	ID             int    `json:"id"`
	Gender         string `json:"gender"`
	DailyEggNumber int    `json:"daily_egg_number"`
}

type MonthlyIncome4 struct {
	MonthlyEggNumber int `json:"monthly_egg_number"`
	EggPrice         int `json:"egg_price"`
	MonthlyMilk      int `json:"monthly_milk_lt"`
	MilkPrice        int `json:"milk_price_lt"`
	MonthlyMeat      int `json:"monthly_meat_kg"`
	MeatPrice        int `json:"meat_price_kg"`
}

type MonthlyExpense4 struct {
	FeedPrice int `json:"feed_price"`
	VetPrice  int `json:"vet_price"`
}

func TakeInfoFarm() {
	// Cattle verilerini al
	if err := fetchAndPrintData("http://localhost:3000/cattle", &[]Cattle4{}); err != nil {
		fmt.Println("Error fetching cattle data:", err)
	}

	// Chicken verilerini al
	if err := fetchAndPrintData("http://localhost:3000/chickens", &[]Chicken4{}); err != nil {
		fmt.Println("Error fetching chicken data:", err)
	}

	// Monthly Income verilerini al
	if err := fetchAndPrintMonthlyIncome("http://localhost:3000/monthly_income"); err != nil {
		fmt.Println("Error fetching monthly income data:", err)
	}

	// Monthly Expense verilerini al
	if err := fetchAndPrintMonthlyExpense("http://localhost:3000/monthly_expense"); err != nil {
		fmt.Println("Error fetching monthly expense data:", err)
	}
}

func fetchAndPrintData(url string, data interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data from %s: %w", url, err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(bodyBytes, data); err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}

	switch v := data.(type) {
	case *[]Cattle4:
		for _, c := range *v {
			fmt.Printf("Name: %s\n", c.Name)
			fmt.Printf("Birthdate: %s\n", c.Birthdate)
			fmt.Printf("Deathdate: %s\n", c.Deathdate)
			fmt.Printf("Ear Number: %d\n", c.EarNumber)
			fmt.Printf("Gender: %s\n", c.Gender)
			fmt.Printf("Mother ID: %d\n", c.MotherID)
			fmt.Printf("Father ID: %d\n", c.FatherID)
			fmt.Printf("Race: %s\n", c.Race)
			fmt.Printf("Color: %s\n", c.Color)
			fmt.Printf("Height: %.2f\n", c.Height)
			fmt.Printf("Chest Size: %v\n", c.ChestSize)
			fmt.Printf("Birth Number: %d\n", c.BirthNumber)
			fmt.Println("-----")
		}
	case *[]Chicken4:
		for _, c := range *v {
			fmt.Printf("Gender: %s\n", c.Gender)
			fmt.Printf("Daily Egg Number: %d\n", c.DailyEggNumber)
			fmt.Println("-----")
		}
	}

	return nil
}

func fetchAndPrintMonthlyIncome(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching monthly income data from %s: %w", url, err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var incomeArr []MonthlyIncome4
	if err := json.Unmarshal(bodyBytes, &incomeArr); err != nil {
		return fmt.Errorf("error unmarshalling monthly income data: %w", err)
	}

	if len(incomeArr) > 0 {
		income := incomeArr[0]
		totalIncome := (income.MonthlyEggNumber * income.EggPrice) +
			(income.MonthlyMeat * income.MeatPrice) +
			(income.MilkPrice * income.MonthlyMilk)

		fmt.Println("Total Monthly Income:", totalIncome)
	} else {
		fmt.Println("No income data available.")
	}

	return nil
}

func fetchAndPrintMonthlyExpense(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching monthly expense data from %s: %w", url, err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var expenseArr []MonthlyExpense4
	if err := json.Unmarshal(bodyBytes, &expenseArr); err != nil {
		return fmt.Errorf("error unmarshalling monthly expense data: %w", err)
	}

	if len(expenseArr) > 0 {
		expense := expenseArr[0]
		totalExpenses := expense.FeedPrice + expense.VetPrice

		fmt.Println("Total Monthly Expenses:", totalExpenses)
	} else {
		fmt.Println("No expense data available.")
	}

	return nil
}
