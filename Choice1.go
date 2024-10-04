package Project1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Cattle1 struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Birthdate    string    `json:"birthdate"`
	Deathdate    string    `json:"deathdate"`
	EarNumber    int       `json:"ear_number"`
	Gender       string    `json:"gender"`
	MotherID     int       `json:"motherId"`
	FatherID     int       `json:"fatherId"`
	Race         string    `json:"race"`
	Color        string    `json:"color"`
	ChestSize    []float64 `json:"chest_size"`
	BirthNumber  int       `json:"birth_number"`
	CauseOfDeath string    `json:"cause_of_death"`
}

type Chicken1 struct {
	ID             int    `json:"id"`
	Gender         string `json:"gender"`
	DailyEggNumber int    `json:"daily_egg_number"`
}

type Vet1 struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PhoneNumber   int    `json:"phone_number"`
	Location      string `json:"location"`
	Expensiveness bool   `json:"expensiveness"`
}

type Feed1 struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Butcher1 struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PhoneNumber   int    `json:"phone_number"`
	AgreementRate bool   `json:"agreement_rate"`
	PriceOfMeat   int    `json:"price_of_meat"`
}

type FeedSeller1 struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phone_number"`
	Price       int    `json:"price"`
}

type MonthlyIncome1 struct {
	MonthlyEggNumber int `json:"monthly_egg_number"`
	EggPrice         int `json:"egg_price"`
	MonthlyMilk      int `json:"monthly_milk_lt"`
	MilkPrice        int `json:"milk_price_lt"`
	MonthlyMeat      int `json:"monthly_meat_kg"`
	MeatPrice        int `json:"meat_price_kg"`
}

type MonthlyExpense1 struct {
	FeedPrice int `json:"feed_price"`
	VetPrice  int `json:"vet_price"`
}

type Database struct {
	Cattle         []Cattle1         `json:"cattle"`
	Chickens       []Chicken1        `json:"chickens"`
	Vets           []Vet1            `json:"vets"`
	Feeds          []Feed1           `json:"feeds"`
	Butchers       []Butcher1        `json:"butchers"`
	FeedSellers    []FeedSeller1     `json:"feed_sellers"`
	MonthlyIncome  []MonthlyIncome1  `json:"monthly_income"`
	MonthlyExpense []MonthlyExpense1 `json:"monthly_expense"`
}

var database1 = Database{
	Cattle: []Cattle1{
		{
			ID:           1,
			Name:         "fatma",
			Birthdate:    "20201010",
			Deathdate:    "0",
			EarNumber:    5,
			Gender:       "male",
			MotherID:     1,
			FatherID:     1,
			Race:         "holstein",
			Color:        "black",
			ChestSize:    []float64{1.75, 1.45},
			BirthNumber:  5,
			CauseOfDeath: "",
		},
	},
	Chickens: []Chicken1{
		{
			ID:             1,
			Gender:         "male",
			DailyEggNumber: 2,
		},
	},
	Vets: []Vet1{
		{
			ID:            1,
			Name:          "Halil",
			PhoneNumber:   5531757714,
			Location:      "Karabogurtlen",
			Expensiveness: true,
		},
	},
	Butchers: []Butcher1{
		{
			ID:            1,
			Name:          "Halil",
			PhoneNumber:   5531757714,
			AgreementRate: true,
			PriceOfMeat:   350,
		},
	},
	Feeds: []Feed1{
		{
			ID:    1,
			Name:  "Cp",
			Price: 250,
		},
	},
	FeedSellers: []FeedSeller1{
		{
			ID:          1,
			Name:        "rt",
			PhoneNumber: 5531757714,
			Price:       555,
		},
	},
	MonthlyIncome: []MonthlyIncome1{
		{
			MonthlyEggNumber: 5,
			EggPrice:         5,
			MonthlyMilk:      5,
			MilkPrice:        5,
			MonthlyMeat:      5,
			MeatPrice:        5,
		},
	},
	MonthlyExpense: []MonthlyExpense1{
		{
			FeedPrice: 5,
			VetPrice:  5,
		},
	},
}

func getNextID(sliceLength int) int {
	return sliceLength + 1
}

func sendDataToAPI(data interface{}, endpoint string) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending data to API:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Data successfully sent to API.")
	} else {
		fmt.Println("Failed to send data to API. Status code:", resp.StatusCode)
	}
}

func AddNewInformation() {
	for {
		fmt.Println("Please choose a category to add a record:")
		fmt.Println("1) Cattle")
		fmt.Println("2) Chickens")
		fmt.Println("3) Vets")
		fmt.Println("4) Feed")
		fmt.Println("5) Butcher")
		fmt.Println("6) Feed Seller")
		fmt.Println("7) Monthly Income")
		fmt.Println("8) Monthly Expense")

		var category int
		_, err := fmt.Scan(&category)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 8.")
			continue
		}

		switch category {
		case 1:
			var newCattle Cattle1
			newCattle.ID = getNextID(len(database1.Cattle))
			fmt.Println("Enter Name:")
			fmt.Scanln(&newCattle.Name)
			fmt.Println("Enter Birthdate (YYYYMMDD):")
			fmt.Scanln(&newCattle.Birthdate)
			fmt.Println("Enter Ear Number:")
			_, err := fmt.Scan(&newCattle.EarNumber)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Gender:")
			fmt.Scanln(&newCattle.Gender)
			fmt.Println("Enter Mother ID:")
			_, err = fmt.Scan(&newCattle.MotherID)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Father ID:")
			_, err = fmt.Scan(&newCattle.FatherID)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Race:")
			fmt.Scanln(&newCattle.Race)
			fmt.Println("Enter Color:")
			fmt.Scanln(&newCattle.Color)
			fmt.Println("Enter Chest Size (two float values separated by space):")
			_, err = fmt.Scan(&newCattle.ChestSize[0], &newCattle.ChestSize[1])
			if err != nil {
				fmt.Println("Invalid input. Please enter two valid float values.")
				continue
			}
			fmt.Println("Enter Birth Number:")
			_, err = fmt.Scan(&newCattle.BirthNumber)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Cause of Death:")
			fmt.Scanln(&newCattle.CauseOfDeath)

			database1.Cattle = append(database1.Cattle, newCattle)
			fmt.Println("New cattle added.")

			// Send data to API
			sendDataToAPI(newCattle, "http://localhost:3000/cattle")

		case 2:
			var newChicken Chicken1
			newChicken.ID = getNextID(len(database1.Chickens))
			fmt.Println("Enter Gender:")
			fmt.Scanln(&newChicken.Gender)
			fmt.Println("Enter Daily Egg Number:")
			_, err := fmt.Scan(&newChicken.DailyEggNumber)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			database1.Chickens = append(database1.Chickens, newChicken)
			fmt.Println("New chicken added.")

			// Send data to API
			sendDataToAPI(newChicken, "http://localhost:3000/chickens")

		case 3:
			var newVet Vet1
			newVet.ID = getNextID(len(database1.Vets))
			fmt.Println("Enter Name:")
			fmt.Scanln(&newVet.Name)
			fmt.Println("Enter Phone Number:")
			_, err := fmt.Scan(&newVet.PhoneNumber)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Location:")
			fmt.Scanln(&newVet.Location)
			fmt.Println("Is the vet expensive? (true/false):")
			var expensivenessStr string
			fmt.Scanln(&expensivenessStr)
			newVet.Expensiveness, err = strconv.ParseBool(expensivenessStr)
			if err != nil {
				fmt.Println("Invalid input. Please enter 'true' or 'false'.")
				continue
			}

			database1.Vets = append(database1.Vets, newVet)
			fmt.Println("New vet added.")

			sendDataToAPI(newVet, "http://localhost:3000/vets")

		case 4:
			var newFeed Feed1
			newFeed.ID = getNextID(len(database1.Feeds))
			fmt.Println("Enter Name:")
			fmt.Scanln(&newFeed.Name)
			fmt.Println("Enter Price:")
			_, err := fmt.Scan(&newFeed.Price)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			database1.Feeds = append(database1.Feeds, newFeed)
			fmt.Println("New feed added.")

			sendDataToAPI(newFeed, "http://localhost:3000/feeds")

		case 5:
			var newButcher Butcher1
			newButcher.ID = getNextID(len(database1.Butchers))
			fmt.Println("Enter Name:")
			fmt.Scanln(&newButcher.Name)
			fmt.Println("Enter Phone Number:")
			_, err := fmt.Scan(&newButcher.PhoneNumber)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Does the butcher have an agreement rate? (true/false):")
			var agreementRateStr string
			fmt.Scanln(&agreementRateStr)
			newButcher.AgreementRate, err = strconv.ParseBool(agreementRateStr)
			if err != nil {
				fmt.Println("Invalid input. Please enter 'true' or 'false'.")
				continue
			}
			fmt.Println("Enter Price of Meat:")
			_, err = fmt.Scan(&newButcher.PriceOfMeat)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			database1.Butchers = append(database1.Butchers, newButcher)
			fmt.Println("New butcher added.")

			sendDataToAPI(newButcher, "http://localhost:3000/butchers")

		case 6:
			var newFeedSeller FeedSeller1
			newFeedSeller.ID = getNextID(len(database1.FeedSellers))
			fmt.Println("Enter Name:")
			fmt.Scanln(&newFeedSeller.Name)
			fmt.Println("Enter Phone Number:")
			_, err := fmt.Scan(&newFeedSeller.PhoneNumber)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Price:")
			_, err = fmt.Scan(&newFeedSeller.Price)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			database1.FeedSellers = append(database1.FeedSellers, newFeedSeller)
			fmt.Println("New feed seller added.")

			sendDataToAPI(newFeedSeller, "http://localhost:3000/feed_sellers")

		case 7:
			var newIncome MonthlyIncome1
			fmt.Println("Enter Monthly Egg Number:")
			_, err := fmt.Scan(&newIncome.MonthlyEggNumber)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Egg Price:")
			_, err = fmt.Scan(&newIncome.EggPrice)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Monthly Milk (liters):")
			_, err = fmt.Scan(&newIncome.MonthlyMilk)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Milk Price (per liter):")
			_, err = fmt.Scan(&newIncome.MilkPrice)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Monthly Meat (kg):")
			_, err = fmt.Scan(&newIncome.MonthlyMeat)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Meat Price (per kg):")
			_, err = fmt.Scan(&newIncome.MeatPrice)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			database1.MonthlyIncome = append(database1.MonthlyIncome, newIncome)
			fmt.Println("New monthly income record added.")

			sendDataToAPI(newIncome, "http://localhost:3000/monthly_income")

		case 8:
			var newExpense MonthlyExpense1
			fmt.Println("Enter Feed Price:")
			_, err := fmt.Scan(&newExpense.FeedPrice)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			fmt.Println("Enter Vet Price:")
			_, err = fmt.Scan(&newExpense.VetPrice)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			database1.MonthlyExpense = append(database1.MonthlyExpense, newExpense)
			fmt.Println("New monthly expense record added.")

			sendDataToAPI(newExpense, "http://localhost:3000/monthly_expense")

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 8.")
		}
	}
}
