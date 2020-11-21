package main

import (
	"fmt"
	"github.com/gopherpay/payment/objectoriented"
	"strings"
)

func main() {
	const amount = 500

	fmt.Println("Paying with cash")
	cash := &objectoriented.Cash{}
	cash.ProcessPayment(amount)
	fmt.Println(strings.Repeat("*", 10) + "\n\n")

	credit := objectoriented.CreateCreditAccount(
		"Zahid Hasan",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	fmt.Println("Paying with credit card")
	fmt.Printf("Initial balance: $%.2f\n", credit.AvailableCredit())
	credit.ProcessPayment(amount)
	fmt.Printf("Balance now: $%2.f\n", credit.AvailableCredit())
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	checking := objectoriented.CreateCheckingAccount(
		"Zahid Hasan",
		"01010011",
		"0551005115100")

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance())
	checking.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f\n", checking.Balance())

	fmt.Println("Not enough in the account. We can do nothing")
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")
}
