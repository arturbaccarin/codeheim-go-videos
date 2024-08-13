/*
Interfaces: implements polymorphism by treating
objects of different types uniformly. (Polymorphism and decoupling)
*/
package interfaces

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Payment of %f processed using card %s", amount, cc.CardNumber)
}

type PayPal struct {
	Email string
}

func (p PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Payment of %f processed using PayPal account %s", amount, p.Email)
}

func process(payment PaymentProcessor, amount float64) {
	fmt.Println(payment.ProcessPayment(amount))
}

func main() {
	cc := CreditCard{CardNumber: "1234-5678-9012-3456"}
	pp := PayPal{Email: "a@b.com"}

	process(cc, 100.95)
	process(pp, 170.95)
}
