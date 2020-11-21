package main

import (
	"fmt"
	"github.com/gopherpay/payment/procedural"
	"strings"
)

func main() {
	const amount = 500

	fmt.Println("Paying with cash")
	procedural.PayWithCash(amount)
	fmt.Println(strings.Repeat("*", 10) + "\n\n")

	credit := &procedural.CreditCard{
		OwnerName: "Zahid Hasan",
		CardNumber: "1111-2222-3333-4444", // sample card number
		ExpirationMonth: 5, // month in 1 to 12 format
		ExpirationYear: 2021,
		SecurityCode: 123,
		AvailableCredit: 5000, // a clear disadvantage of procedural programming
		// we can manipulate the data here
	}

	fmt.Println("Paying with credit card")
	fmt.Printf("Initial balance: $%.2f\n", credit.AvailableCredit)
	procedural.PayWithCredit(credit, amount)
	fmt.Printf("Balance now: $%2.f\n", credit.AvailableCredit)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	checking := &procedural.CheckingAccount{
		AccountOwner: "Zahid Hasan",
		RoutingNumber: "01010011",
		AccountNumber: "0551005115100",
		Balance: 250, // same disadvantage in procedural programming
	}

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance)
	procedural.PayWithCheck(checking, amount)
	fmt.Printf("Balance now: $%2.f\n", checking.Balance)

	fmt.Println("Not enough in the account. We can fix that")
	fmt.Println("Changing account balance")
	checking.Balance = 5000

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance)
	procedural.PayWithCheck(checking, amount)
	fmt.Printf("Balance now: $%2.f\n", checking.Balance)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")
}