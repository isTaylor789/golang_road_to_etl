package main

import (
	"fmt"
)

// PaymentProcessor define la interfaz común para todos los métodos de pago
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// CreditCard es una implementación de PaymentProcessor
type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Procesando un pago de $%.2f con tarjeta de crédito: %s", amount, cc.CardNumber)
}

// BankTransfer es otra implementación de PaymentProcessor
type BankTransfer struct {
	AccountNumber string
}

func (bt BankTransfer) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Procesando una transferencia bancaria de $%.2f desde la cuenta: %s", amount, bt.AccountNumber)
}

// PayPal es otra implementación de PaymentProcessor
type PayPal struct {
	Email string
}

func (pp PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Procesando un pago de $%.2f a través de PayPal con el email: %s", amount, pp.Email)
}

// main muestra el uso del polimorfismo
func main() {
	// Crear instancias de los métodos de pago
	creditCard := CreditCard{CardNumber: "1234-5678-9101-1121"}
	bankTransfer := BankTransfer{AccountNumber: "9876543210"}
	payPal := PayPal{Email: "user@example.com"}

	// Lista de procesadores de pago
	payments := []PaymentProcessor{
		creditCard,
		bankTransfer,
		payPal,
	}

	// Procesar pagos de forma polimórfica
	for _, payment := range payments {
		fmt.Println(payment.ProcessPayment(100.00))
	}
}
