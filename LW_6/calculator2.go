package main

import (
	"fmt"
	"net/url"
)

type Task2Calculations struct{}

// приймає масиви значень та виконує розрахунки
func (tc *Task2Calculations) Calculate(quantity, power, usageCoeff, tgPhi, voltage []float64) map[string]float64 {
	numerator, denominator, powerSum, powerSquaredSum, reactivePowerSum, totalVoltage := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	voltageCount := 0

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
	}

	// щоб уникнути ділення на 0
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
	Ip := Pp / Uh

	return map[string]float64{
		"nPh":             denominator,
		"nPhKv":           numerator,
		"Qp":              Qp,
		"powerSquaredSum": powerSquaredSum,
		"Ip":              Ip,
	}
}

// зчитує значення, обробляє їх і викликає `Calculate`
func calculateTask2(values url.Values) map[string]interface{} {
	// зчитуємо параметри з форми
	quantityStrings := values["quantityController[]"]
	powerStrings := values["powerController[]"]
	usageCoeffStrings := values["usageCoeffController[]"]
	tgPhiStrings := values["tgPhiController[]"]
	voltageStrings := values["voltageController[]"]

	// претворення у float64
	quantity, power, usageCoeff, tgPhi, voltage := convertToFloatSlice(quantityStrings), convertToFloatSlice(powerStrings), convertToFloatSlice(usageCoeffStrings), convertToFloatSlice(tgPhiStrings), convertToFloatSlice(voltageStrings)

	calculator := Task2Calculations{}
	results := calculator.Calculate(quantity, power, usageCoeff, tgPhi, voltage)

	// форматований вивід
	return map[string]interface{}{
		"nPh":             fmt.Sprintf("%.2f", results["nPh"]),
		"nPhKv":           fmt.Sprintf("%.2f", results["nPhKv"]),
		"Qp":              fmt.Sprintf("%.2f", results["Qp"]),
		"powerSquaredSum": fmt.Sprintf("%.2f", results["powerSquaredSum"]),
		"Ip":              fmt.Sprintf("%.2f", results["Ip"]),
	}
}
