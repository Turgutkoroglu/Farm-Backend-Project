package Project1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Cattle3 struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Color        string    `gorm:"column:color" json:"color"`
	BirthNumber  int       `gorm:"column:birth_number" json:"birth_number"`
	CauseOfDeath string    `gorm:"column:cause_of_death" json:"cause_of_death"`
	Deathdate    string    `json:"deathdate"`
	ChestSize    []float64 `json:"chest_size"`
}

type Chicken3 struct {
	ID             int    `gorm:"primaryKey" json:"id"`
	Gender         string `json:"gender"`
	DailyEggNumber int    `json:"daily_egg_number"`
}

type Vet3 struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	PhoneNumber   int    `json:"phone_number"`
	Location      string `json:"location"`
	Expensiveness bool   `json:"expensiveness"`
}

type Feed3 struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Price int    `json:"price"`
	Name  string `json:"name"`
}

type Butcher3 struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	PhoneNumber   int    `json:"phone_number"`
	AgreementRate bool   `json:"agreement_rate"`
	PriceOfMeat   int    `json:"price_of_meat"`
}

type FeedSeller3 struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Brand       string `json:"brand"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phone_number"`
	Price       int    `json:"price"`
}

type MonthlyIncome3 struct {
	MonthlyEggNumber int `json:"monthly_egg_number"`
	EggPrice         int `json:"egg_price"`
	MonthlyMilk      int `json:"monthly_milk_lt"`
	MilkPrice        int `json:"milk_price_lt"`
	MonthlyMeat      int `json:"monthly_meat_kg"`
	MeatPrice        int `json:"meat_price_kg"`
}

type MonthlyExpense3 struct {
	FeedPrice int `json:"feed_price"`
	VetPrice  int `json:"vet_price"`
}

func UpdateInformation() {
	db, err := gorm.Open(sqlite.Open("farm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	}

	db.AutoMigrate(&Cattle3{}, &Chicken3{}, &Vet3{}, &Feed3{}, &Butcher3{}, &FeedSeller3{}, &MonthlyIncome3{}, &MonthlyExpense3{})

	var option int
	fmt.Println("Which data would you like to update?")
	fmt.Println("1: Cattle\n2: Chicken\n3: Vet\n4: Feed\n5: Butcher\n6: FeedSeller\n7: MonthlyIncome\n8: MonthlyExpense")
	fmt.Scan(&option)

	switch option {
	case 1:
		updateCattle(db)
	case 2:
		updateChicken(db)
	case 3:
		updateVet(db)
	case 4:
		updateFeed(db)
	case 5:
		updateButcher(db)
	case 6:
		updateFeedSeller(db)
	case 7:
		updateMonthlyIncome(db)
	case 8:
		updateMonthlyExpense(db)
	default:
		fmt.Println("Invalid option selected.")
	}
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func updateCattle(db *gorm.DB) {
	var id int
	fmt.Print("Enter the ID of the cattle to update: ")
	fmt.Scan(&id)

	var cattle Cattle3
	if err := db.First(&cattle, id).Error; err != nil {
		fmt.Println("Cattle not found:", err)
		return
	}

	cattle.Color = getInput("Enter new color: ")
	cattle.BirthNumber, _ = strconv.Atoi(getInput("Enter new birth number: "))
	cattle.CauseOfDeath = getInput("Enter new cause of death: ")
	cattle.Deathdate = getInput("Enter new deathdate: ")

	db.Save(&cattle)
	fmt.Println("Cattle updated successfully.")
}

func updateChicken(db *gorm.DB) {
	var id int
	fmt.Print("Enter the ID of the chicken to update: ")
	fmt.Scan(&id)

	var chicken Chicken3
	if err := db.First(&chicken, id).Error; err != nil {
		fmt.Println("Chicken not found:", err)
		return
	}

	chicken.Gender = getInput("Enter new gender: ")
	chicken.DailyEggNumber, _ = strconv.Atoi(getInput("Enter new daily egg number: "))

	db.Save(&chicken)
	fmt.Println("Chicken updated successfully.")
}

func updateVet(db *gorm.DB) {
	var id int
	fmt.Print("Enter the ID of the vet to update: ")
	fmt.Scan(&id)

	var vet Vet3
	if err := db.First(&vet, id).Error; err != nil {
		fmt.Println("Vet not found:", err)
		return
	}

	vet.PhoneNumber, _ = strconv.Atoi(getInput("Enter new phone number: "))
	vet.Location = getInput("Enter new location: ")
	expensiveness, _ := strconv.ParseBool(getInput("Is the vet expensive? (true/false): "))
	vet.Expensiveness = expensiveness

	db.Save(&vet)
	fmt.Println("Vet updated successfully.")
}

func updateFeed(db *gorm.DB) {
	var id int
	fmt.Print("Enter the ID of the feed to update: ")
	fmt.Scan(&id)

	var feed Feed3
	if err := db.First(&feed, id).Error; err != nil {
		fmt.Println("Feed not found:", err)
		return
	}

	feed.Price, _ = strconv.Atoi(getInput("Enter new price: "))
	feed.Name = getInput("Enter new name: ")

	db.Save(&feed)
	fmt.Println("Feed updated successfully.")
}

func updateButcher(db *gorm.DB) {
	var id int
	fmt.Print("Enter the ID of the butcher to update: ")
	fmt.Scan(&id)

	var butcher Butcher3
	if err := db.First(&butcher, id).Error; err != nil {
		fmt.Println("Butcher not found:", err)
		return
	}

	butcher.Name = getInput("Enter new name: ")
	butcher.PhoneNumber, _ = strconv.Atoi(getInput("Enter new phone number: "))
	agreementRate, _ := strconv.ParseBool(getInput("Is there an agreement? (true/false): "))
	butcher.AgreementRate = agreementRate
	butcher.PriceOfMeat, _ = strconv.Atoi(getInput("Enter new price of meat: "))

	db.Save(&butcher)
	fmt.Println("Butcher updated successfully.")
}

func updateFeedSeller(db *gorm.DB) {
	var id int
	fmt.Print("Enter the ID of the feed seller to update: ")
	fmt.Scan(&id)

	var feedSeller FeedSeller3
	if err := db.First(&feedSeller, id).Error; err != nil {
		fmt.Println("Feed seller not found:", err)
		return
	}

	feedSeller.Brand = getInput("Enter new brand: ")
	feedSeller.Name = getInput("Enter new name: ")
	feedSeller.PhoneNumber, _ = strconv.Atoi(getInput("Enter new phone number: "))
	feedSeller.Price, _ = strconv.Atoi(getInput("Enter new price: "))

	db.Save(&feedSeller)
	fmt.Println("Feed seller updated successfully.")
}

func updateMonthlyIncome(db *gorm.DB) {
	var income MonthlyIncome3

	income.MonthlyEggNumber, _ = strconv.Atoi(getInput("Enter new monthly egg number: "))
	income.EggPrice, _ = strconv.Atoi(getInput("Enter new egg price: "))
	income.MonthlyMilk, _ = strconv.Atoi(getInput("Enter new monthly milk (liters): "))
	income.MilkPrice, _ = strconv.Atoi(getInput("Enter new milk price (per liter): "))
	income.MonthlyMeat, _ = strconv.Atoi(getInput("Enter new monthly meat (kg): "))
	income.MeatPrice, _ = strconv.Atoi(getInput("Enter new meat price (per kg): "))

	db.Save(&income)
	fmt.Println("Monthly income updated successfully.")
}

func updateMonthlyExpense(db *gorm.DB) {
	var expense MonthlyExpense3

	expense.FeedPrice, _ = strconv.Atoi(getInput("Enter new feed price: "))
	expense.VetPrice, _ = strconv.Atoi(getInput("Enter new vet price: "))

	db.Save(&expense)
	fmt.Println("Monthly expense updated successfully.")
}
