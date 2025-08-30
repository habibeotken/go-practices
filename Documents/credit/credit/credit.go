package credit

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}
type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 60
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 36
}
func MonthlyPayment(credits []CreditCalculator) float64 {
	total := 0.0
	for _, c := range credits {
		total += c.Calculate()
	}
	return total
}
func Demo2() {
	credit1 := Mortgage{
		creditPaymentTotal: 10000000,
		address:            "123 Main St",
		rate:               0.05,
	}
	credit2 := Car{
		creditPaymentTotal: 30000000,
		carInfo:            "Range Rover",
		rate:               0.04,
	}
	credits := []CreditCalculator{credit1, credit2}
	fmt.Println("Total Monthly Payment:", MonthlyPayment(credits))
}
