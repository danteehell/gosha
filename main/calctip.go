package main

import "fmt"

func calcTip(billAmount float64, tipPercentage float64, numberOfPeople int) {
	// я бы добавила проверку на nil и подстановку дефолтных значений, но лень
	// Расчёт
	tipAmount := billAmount * tipPercentage / 100.0
	totalAmount := billAmount + tipAmount
	perPersonAmount := totalAmount / float64(numberOfPeople)

	fmt.Printf("Сумма счёта: %.2f\nПроцент чаевых: %.2f\n"+
		"Сколько человек участвовало в обеде: %d\nОбщий счёт: %.2f\n"+
		"Процент чаевых: %.1f%%\nС каждого: %.2f",
		billAmount, tipPercentage, numberOfPeople, totalAmount, tipPercentage, perPersonAmount)
}
