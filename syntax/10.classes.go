package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Classes

// 1. Basic Struct with Constructor and Instance Variables
type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func NewPersonWithDefault(name string) *Person {
	return NewPerson(name, 36)
}

func (p *Person) introduce() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

// 2. Struct with Package-level Variables and Functions
type Employee struct {
	name string
	age  int
}

var (
	company        = "TechCorp"
	employeeCount  = 0
)

func NewEmployee(name string, age int) *Employee {
	employeeCount++
	return &Employee{name: name, age: age}
}

func (e *Employee) introduce() {
	fmt.Printf("Hello, I'm %s, working at %s\n", e.name, company)
}

func getCompanyInfo() string {
	return fmt.Sprintf("Company: %s, Total Employees: %d", company, employeeCount)
}

// 3. Struct with Private Variables and Methods
type BankAccount struct {
	accountHolder string
	balance       float64
	accountNumber string
}

func NewBankAccount(accountHolder string, initialBalance float64) *BankAccount {
	return &BankAccount{
		accountHolder: accountHolder,
		balance:       initialBalance,
		accountNumber: generateAccountNumber(),
	}
}

func generateAccountNumber() string {
	return "ACC" + strconv.Itoa(10000+rand.Intn(90000))
}

func (ba *BankAccount) deposit(amount float64) {
	if amount > 0 {
		ba.balance += amount
		fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, ba.balance)
	} else {
		fmt.Println("Invalid deposit amount")
	}
}

func (ba *BankAccount) getBalance() float64 {
	return ba.balance
}

func (ba *BankAccount) introduce() {
	fmt.Printf("Account holder: %s, Balance: $%.2f\n", ba.accountHolder, ba.balance)
}

// 4. Struct with Embedding (Inheritance-like)
type Student struct {
	Person
	grade     string
	studentId string
}

func NewStudent(name string, age int, grade, studentId string) *Student {
	return &Student{
		Person:    *NewPerson(name, age),
		grade:     grade,
		studentId: studentId,
	}
}

func (s *Student) introduce() {
	fmt.Printf("Hello, I'm %s, a student with ID %s and grade %s\n", 
		s.name, s.studentId, s.grade)
}

func (s *Student) study() {
	fmt.Printf("%s is studying hard!\n", s.name)
}

// 5. Struct with Property-like behavior
type Temperature struct {
	celsius float64
}

func NewTemperature(celsius float64) *Temperature {
	t := &Temperature{}
	t.setCelsius(celsius)
	return t
}

func (t *Temperature) getCelsius() float64 {
	return t.celsius
}

func (t *Temperature) setCelsius(value float64) {
	if value < -273.15 { // Absolute zero
		panic("Temperature cannot be below absolute zero")
	}
	t.celsius = value
}

func (t *Temperature) getFahrenheit() float64 {
	return (t.celsius * 9 / 5) + 32
}

func main() {
	fmt.Println("=== Testing Basic Person Struct ===")
	person1 := NewPerson("John", 36)
	person1.introduce()

	fmt.Println("\n=== Testing Employee Struct ===")
	emp1 := NewEmployee("Alice", 28)
	emp2 := NewEmployee("Bob", 32)
	emp1.introduce()
	emp2.introduce()
	fmt.Println(getCompanyInfo())

	fmt.Println("\n=== Testing BankAccount Struct ===")
	account := NewBankAccount("John Doe", 1000)
	account.deposit(500)
	account.introduce()
	fmt.Printf("Current balance: $%.2f\n", account.getBalance())

	fmt.Println("\n=== Testing Student Struct (Embedding) ===")
	student := NewStudent("Jane", 20, "A", "STU123")
	student.introduce()
	student.study()

	fmt.Println("\n=== Testing Temperature Struct (Properties) ===")
	temp := NewTemperature(25)
	fmt.Printf("Celsius: %.1f°C\n", temp.getCelsius())
	fmt.Printf("Fahrenheit: %.1f°F\n", temp.getFahrenheit())
	temp.setCelsius(30) // Using setter
	fmt.Printf("New Celsius: %.1f°C\n", temp.getCelsius())
} 