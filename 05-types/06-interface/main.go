package main

import "fmt"

type Card struct {
	Number      string
	HolderName  string
	ExpiryMonth uint8
	ExpiryYear  uint8
	CVV         string
}

type PaymentStrategy interface {
	PayWithCreditCard(card Card)
}

type AsaasPaymentGw struct {
	BaseUrl string
	ApiKey  string
}

func (gw AsaasPaymentGw) PayWithCreditCard(card Card) {
	fmt.Println("Execute asass payment method")
}

type StripePaymentGw struct {
	BaseUrl string
	ApiKey  string
}

func (gw StripePaymentGw) PayWithCreditCard(card Card) {
	fmt.Println("Execute stripe payment method")
}

func executePayment(strategy PaymentStrategy, card Card) {
	strategy.PayWithCreditCard(card)
}

func main() {
	stripeGw := StripePaymentGw{}
	asassGw := AsaasPaymentGw{}
	card := Card{}

	// interfaces são implementadas implicitamente
	executePayment(stripeGw, card)
	executePayment(asassGw, card)
}
