// https://www.vtbnpf.ru/
// калькулятор выплат по вкладу (Инд пенс план)
package main

import (
	"fmt"
	"math"
)

var (
	initialFee                = 0.0   // первоначальный взнос
	lAccumulation             = 10.0  // years - срок накопления до выхода на пенсию
	mounthlyFee               = 70000 // rubs - ежемесячный взнос
	profitabilityAccumulation = 6.51  // % - прогнозируемая ставка на период накопления
	profitabilityPay          = 4.0   // актуарный % - для расчета доходности на 13 лет, const
	lPayForCalculation        = 13.0  // years - расчётный срок для расчета ежемесячной суммы
	// mounthlyPay               = 97084.0
)

func main() {
	fundsWithIncom := initialFee // доходность в основной период
	iMonths := (int)(lAccumulation * 12)
	i := 0
	for i = 0; i < iMonths; i++ {
		fundsWithIncom = fundsWithIncom*(1+profitabilityAccumulation/12/100) + float64(mounthlyFee)
	}
	fmt.Printf("Накоплено за %.1f лет: %d рублей\n", lAccumulation, int(fundsWithIncom))
	payWithoutIncom := fundsWithIncom / lPayForCalculation / 12
	fmt.Printf("Без индексации ежемесячная выплата за %.1f лет: %d рублей\n", lPayForCalculation, int(payWithoutIncom))
	// payWithIncom := fundsWithIncom / lPayForCalculation / 12 * 1.284
	// fmt.Printf("С индексацией ежемесячная выплата за %.1f лет: %d рублей\n", lPayForCalculation, int(payWithIncom))
	// payWithIncom = fundsWithIncom / lPayForCalculation / 12 * 1.431
	// fmt.Printf("С индексацией ежемесячная выплата за %.1f лет: %d рублей\n", lPayForCalculation, int(payWithIncom))
	// iMonths = (int)(lPayForCalculation * 12)

	// https://www.raiffeisen.ru/wiki/kak-rasschitat-procenty-po-kreditu/
	p := profitabilityPay / 100 / 12
	mounthlyPay := fundsWithIncom * (p + (p / (math.Pow(1+p, lPayForCalculation*12) - 1)))
	fmt.Printf("С индексацией ПОЖИЗНЕННАЯ ежемесячная выплата (за %.1f лет): %d рублей\n", lPayForCalculation, int(mounthlyPay))
}

// go run incom.go
