package Project1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Cattle2 struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Birthdate    string    `json:"birthdate"`
	Deathdate    string    `json:"deathdate"`
	EarNumber    int       `json:"ear_number"`
	Gender       string    `json:"gender"`
	MotherID     int       `json:"mother_id"`
	FatherID     int       `json:"father_id"`
	Race         string    `json:"race"`
	Color        string    `json:"color"`
	ChestSize    []float64 `json:"chest_size"`
	BirthNumber  int       `json:"birth_number"`
	CauseOfDeath string    `json:"cause_of_death"`
}

type Chicken2 struct {
	ID             int    `json:"id"`
	Gender         string `json:"gender"`
	DailyEggNumber int    `json:"daily_egg_number"`
}

type Vet2 struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PhoneNumber   int    `json:"phone_number"`
	Location      string `json:"location"`
	Expensiveness bool   `json:"expensiveness"`
}

type Feed2 struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Butcher2 struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PhoneNumber   int    `json:"phone_number"`
	AgreementRate bool   `json:"agreement_rate"`
	PriceOfMeat   int    `json:"price_of_meat"`
}

type FeedSeller2 struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phone_number"`
	Price       int    `json:"price"`
}

type MonthlyIncome2 struct {
	MonthlyEggNumber int `json:"monthly_egg_number"`
	EggPrice         int `json:"egg_price"`
	MonthlyMilk      int `json:"monthly_milk"`
	MilkPrice        int `json:"milk_price"`
	MonthlyMeat      int `json:"monthly_meat"`
	MeatPrice        int `json:"meat_price"`
}

type MonthlyExpense2 struct {
	FeedPrice int `json:"feed_price"`
	VetPrice  int `json:"vet_price"`
}

// List functions
func ListCattle() ([]Cattle2, error) {
	response, err := http.Get("http://localhost:3000/cattle")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var cattle []Cattle2
	err = json.Unmarshal(bodyBytes, &cattle)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return cattle, nil
}
func ListChickens() ([]Chicken2, error) {
	response, err := http.Get("http://localhost:3000/chickens")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var chickens []Chicken2
	err = json.Unmarshal(bodyBytes, &chickens)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return chickens, nil
}

func ListVets() ([]Vet2, error) {
	response, err := http.Get("http://localhost:3000/vets")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var vets []Vet2
	err = json.Unmarshal(bodyBytes, &vets)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return vets, nil
}

func ListFeeds() ([]Feed2, error) {
	response, err := http.Get("http://localhost:3000/feeds")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var feeds []Feed2
	err = json.Unmarshal(bodyBytes, &feeds)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return feeds, nil
}
func ListButchers() ([]Butcher2, error) {
	response, err := http.Get("http://localhost:3000/butchers")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var butchers []Butcher2
	err = json.Unmarshal(bodyBytes, &butchers)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return butchers, nil
}

func ListFeedSellers() ([]FeedSeller2, error) {
	response, err := http.Get("http://localhost:3000/feed_sellers")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var sellers []FeedSeller2
	err = json.Unmarshal(bodyBytes, &sellers)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return sellers, nil
}

func ListMonthlyIncome() ([]MonthlyIncome2, error) {
	response, err := http.Get("http://localhost:3000/monthly_income")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var income []MonthlyIncome2
	err = json.Unmarshal(bodyBytes, &income)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return income, nil
}

func ListMonthlyExpenses() ([]MonthlyExpense2, error) {
	response, err := http.Get("http://localhost:3000/monthly_expense")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var expenses []MonthlyExpense2
	err = json.Unmarshal(bodyBytes, &expenses)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return expenses, nil
}

// Print functions
func PrintCattleInfo() {
	cattle, err := ListCattle()
	if err != nil {
		fmt.Println("Error fetching cattle data:", err)
		return
	}
	for _, c := range cattle {
		fmt.Printf("ID: %d\nName: %s\nBirthdate: %s\nDeathdate: %s\nEarNumber: %d\nGender: %s\nMotherID: %d\nFatherID: %d\nRace: %s\nColor: %s\nChestSize: %v\nBirthNumber: %d\nCauseOfDeath: %s\n\n",
			c.ID, c.Name, c.Birthdate, c.Deathdate, c.EarNumber, c.Gender, c.MotherID, c.FatherID, c.Race, c.Color, c.ChestSize, c.BirthNumber, c.CauseOfDeath)
	}
}

func PrintChickenInfo() {
	chickens, err := ListChickens()
	if err != nil {
		fmt.Println("Error fetching chicken data:", err)
		return
	}
	for _, c := range chickens {
		fmt.Printf("ID: %d\nGender: %s\nDailyEggNumber: %d\n\n", c.ID, c.Gender, c.DailyEggNumber)
	}
}

func PrintVetInfo() {
	vets, err := ListVets()
	if err != nil {
		fmt.Println("Error fetching vet data:", err)
		return
	}
	for _, v := range vets {
		fmt.Printf("ID: %d\nName: %s\nPhoneNumber: %d\nLocation: %s\nExpensiveness: %t\n\n",
			v.ID, v.Name, v.PhoneNumber, v.Location, v.Expensiveness)
	}
}

func PrintFeedInfo() {
	feeds, err := ListFeeds()
	if err != nil {
		fmt.Println("Error fetching feed data:", err)
		return
	}
	for _, f := range feeds {
		fmt.Printf("ID: %d\nName: %s\nPrice: %d\n\n", f.ID, f.Name, f.Price)
	}
}

func PrintButcherInfo() {
	butchers, err := ListButchers()
	if err != nil {
		fmt.Println("Error fetching butcher data:", err)
		return
	}
	for _, b := range butchers {
		fmt.Printf("ID: %d\nName: %s\nPhoneNumber: %d\nAgreementRate: %t\nPriceOfMeat: %d\n\n",
			b.ID, b.Name, b.PhoneNumber, b.AgreementRate, b.PriceOfMeat)
	}
}

func PrintFeedSellerInfo() {
	sellers, err := ListFeedSellers()
	if err != nil {
		fmt.Println("Error fetching feed seller data:", err)
		return
	}
	for _, s := range sellers {
		fmt.Printf("ID: %d\nName: %s\nPhoneNumber: %d\nPrice: %d\n\n", s.ID, s.Name, s.PhoneNumber, s.Price)
	}
}

func PrintMonthlyIncomeInfo() {
	income, err := ListMonthlyIncome()
	if err != nil {
		fmt.Println("Error fetching monthly income data:", err)
		return
	}
	for _, i := range income {
		fmt.Printf("MonthlyEggNumber: %d\nEggPrice: %d\nMonthlyMilk: %d\nMilkPrice: %d\nMonthlyMeat: %d\nMeatPrice: %d\n\n",
			i.MonthlyEggNumber, i.EggPrice, i.MonthlyMilk, i.MilkPrice, i.MonthlyMeat, i.MeatPrice)
	}
}

func PrintMonthlyExpenseInfo() {
	expenses, err := ListMonthlyExpenses()
	if err != nil {
		fmt.Println("Error fetching monthly expense data:", err)
		return
	}
	for _, e := range expenses {
		fmt.Printf("FeedPrice: %d\nVetPrice: %d\n\n", e.FeedPrice, e.VetPrice)
	}
}
