package main

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

type Task1Calculations struct{}

// приймає масиви значень та виконує розрахунки
func (tc *Task1Calculations) Calculate(quantity, power, usageCoeff, tgPhi, voltage []float64) map[string]float64 {
	numerator, denominator, powerSum, powerSquaredSum, reactivePowerSum, totalVoltage := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	voltageCount := 0
	quantityNum := 0.0

	for i := range quantity {
		if voltage[i] > 0 {
			totalVoltage += voltage[i]
			voltageCount++
		}

		numerator += quantity[i] * power[i] * usageCoeff[i]
		denominator += quantity[i] * power[i]
		powerSum += quantity[i] * power[i]
		powerSquaredSum += quantity[i] * power[i] * power[i]
		reactivePowerSum += quantity[i] * power[i] * usageCoeff[i] * tgPhi[i]
		quantityNum += quantity[i]
	}

	//  щоб уникнути ділення на 0
	Uh := 0.0
	if voltageCount > 0 {
		Uh = totalVoltage / float64(voltageCount)
	}
	if Uh == 0.0 {
		return map[string]float64{"error": -1}
	}

	Kv := 0.0
	if denominator > 0 {
		Kv = numerator / denominator
	}
	Ne := 0.0
	if powerSquaredSum > 0 {
		Ne = (powerSum * powerSum) / powerSquaredSum
	}

	Kp := GetActivePowerCoefficient(Kv, Ne)
	Pp := Kp * numerator
	Qp := reactivePowerSum
	Sp := math.Sqrt(Pp*Pp + Qp*Qp)
	Ip := Pp / Uh

	return map[string]float64{
		"Kv":       Kv,
		"Ne":       Ne,
		"Kp":       Kp,
		"Pp":       Pp,
		"Qp":       Qp,
		"Sp":       Sp,
		"Ip":       Ip,
		"quantity": float64(quantityNum),
		"nPh":      denominator,
		"nPhKv":    numerator,
		"nPh2":     powerSquaredSum,
	}
}

// зчитує значення, обробляє їх і викликає Calculate
func calculateTask1(values url.Values) map[string]interface{} {
	// зчитуємо параметри з форми
	quantityStrings := values["quantityController[]"]
	powerStrings := values["powerController[]"]
	usageCoeffStrings := values["usageCoeffController[]"]
	tgPhiStrings := values["tgPhiController[]"]
	voltageStrings := values["voltageController[]"]

	// перетворення у float64
	quantity, power, usageCoeff, tgPhi, voltage := convertToFloatSlice(quantityStrings), convertToFloatSlice(powerStrings), convertToFloatSlice(usageCoeffStrings), convertToFloatSlice(tgPhiStrings), convertToFloatSlice(voltageStrings)

	// виклик розрахунків
	calculator := Task1Calculations{}
	results := calculator.Calculate(quantity, power, usageCoeff, tgPhi, voltage)

	// форматований вивід
	return map[string]interface{}{
		"Kv":       fmt.Sprintf("%.2f", results["Kv"]),
		"Ne":       fmt.Sprintf("%.2f", results["Ne"]),
		"Kp":       fmt.Sprintf("%.2f", results["Kp"]),
		"Pp":       fmt.Sprintf("%.2f", results["Pp"]),
		"Qp":       fmt.Sprintf("%.2f", results["Qp"]),
		"Sp":       fmt.Sprintf("%.2f", results["Sp"]),
		"Ip":       fmt.Sprintf("%.2f", results["Ip"]),
		"quantity": fmt.Sprintf("%.2f", results["quantity"]),
		"nPh":      fmt.Sprintf("%.2f", results["nPh"]),
		"nPhKv":    fmt.Sprintf("%.2f", results["nPhKv"]),
		"nPh2":     fmt.Sprintf("%.2f", results["nPh2"]),
	}
}

// convertToFloatSlice перетворює масив рядків у масив чисел
func convertToFloatSlice(values []string) []float64 {
	var result []float64
	for _, val := range values {
		num := parseFloat(val)
		result = append(result, num)
	}
	return result
}

// parseFloat обробляє конвертацію з рядка у float64
func parseFloat(value string) float64 {
	if value == "" {
		return 0
	}
	value = strings.Replace(value, ",", ".", -1)
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Помилка парсингу:", value, err)
		return 0
	}
	return num
}
