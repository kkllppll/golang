package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Task2Calculations struct{}

// розраховує математичне сподівання аварійного та планового недовідпуску та загальні збитки
func (tc *Task2Calculations) Calculate(omega, tb, p, T, k, priceA, priceP float64) map[string]float64 {
	mWnedA := omega * tb * p * T
	mWnedP := k * p * T
	Mzper := (priceA * mWnedA) + (priceP * mWnedP)

	return map[string]float64{
		"mWnedA": mWnedA,
		"mWnedP": mWnedP,
		"Mzper":  Mzper,
	}
}

func parseFloat(value string) float64 {
	value = strings.Replace(value, ",", ".", -1)
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {

		fmt.Println("Помилка парсингу:", value)
		return 0.0 // або значення за замовчуванням
	}
	return num
}

// отримує значення з параметрів URL та виконує обчислення
func calculateTask2(values url.Values) map[string]interface{} {
	omega := parseFloat(values.Get("omegaController"))
	tb := parseFloat(values.Get("tbController"))
	p := parseFloat(values.Get("pController"))
	T := parseFloat(values.Get("tController"))
	k := parseFloat(values.Get("kController"))
	priceA := parseFloat(values.Get("priceAController"))
	priceP := parseFloat(values.Get("pricePController"))

	calculator := Task2Calculations{}
	results := calculator.Calculate(omega, tb, p, T, k, priceA, priceP)

	//  результати
	return map[string]interface{}{
		"mWnedA": fmt.Sprintf("%.2f", results["mWnedA"]),
		"mWnedP": fmt.Sprintf("%.2f", results["mWnedP"]),
		"Mzper":  fmt.Sprintf("%.2f", results["Mzper"]),
	}
}
