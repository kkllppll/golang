package main

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
)

type Task1Calculations struct{}

func (tc *Task1Calculations) Calculate(omegaValues, recoveryTimeValues []float64, coefSimpleDowntime float64) map[string]float64 {
	omegaSystem := 0.0
	avgRecoveryTime := 0.0

	for i := range omegaValues {
		omegaSystem += omegaValues[i]
	}

	// обчислення середнього часу відновлення
	if omegaSystem > 0 {
		for i := range omegaValues {
			avgRecoveryTime += omegaValues[i] * recoveryTimeValues[i]
		}
		avgRecoveryTime /= omegaSystem
		avgRecoveryTime = math.Round(avgRecoveryTime*10) / 10.0 // округлення
	}

	kAoc := (omegaSystem * avgRecoveryTime) / 8760
	kPlOc := (1.2 * coefSimpleDowntime) / 8760
	omegaDk := 2 * omegaSystem * (kAoc + kPlOc)
	omegaDs := omegaDk + 0.02

	return map[string]float64{
		"omegaSystem":     omegaSystem,
		"avgRecoveryTime": avgRecoveryTime,
		"kAoc":            kAoc,
		"kPlOc":           kPlOc,
		"omegaDk":         omegaDk,
		"omegaDs":         omegaDs,
	}
}

func calculateTask1(values url.Values) map[string]interface{} {
	omegaValues := values["omega[]"]
	recoveryTimeValues := values["recoveryTime[]"]
	coefSimpleDowntimeStr := values.Get("coefSimpleDowntime")

	// перетворення рядкових значень у числа
	var omegaSystem float64
	var avgRecoveryTime float64
	omegas := []float64{}
	recoveryTimes := []float64{}

	for i := range omegaValues {
		omega, err := strconv.ParseFloat(omegaValues[i], 64)
		if err != nil {
			return map[string]interface{}{"error": "Invalid omega value"}
		}

		recoveryTime, err := strconv.ParseFloat(recoveryTimeValues[i], 64)
		if err != nil {
			return map[string]interface{}{"error": "Invalid recovery time value"}
		}

		omegas = append(omegas, omega)
		recoveryTimes = append(recoveryTimes, recoveryTime)
		omegaSystem += omega
	}

	// обчислення середнього часу відновлення
	if omegaSystem > 0 {
		for i := range omegas {
			avgRecoveryTime += omegas[i] * recoveryTimes[i]
		}
		avgRecoveryTime /= omegaSystem
		avgRecoveryTime = math.Round(avgRecoveryTime*10) / 10.0 // округлення
	}

	// коефіцієнти
	kPlMax, err := strconv.ParseFloat(coefSimpleDowntimeStr, 64)
	if err != nil {
		kPlMax = 0.0 // якщо введено некоректне значення
	}

	calculator := Task1Calculations{}
	results := calculator.Calculate(omegas, recoveryTimes, kPlMax)

	return map[string]interface{}{
		"omegaSystem":     results["omegaSystem"],
		"avgRecoveryTime": results["avgRecoveryTime"],
		"kAoc":            fmt.Sprintf("%.5f", results["kAoc"]),
		"kPlOc":           fmt.Sprintf("%.5f", results["kPlOc"]),
		"omegaDk":         fmt.Sprintf("%.5f", results["omegaDk"]),
		"omegaDs":         fmt.Sprintf("%.4f", results["omegaDs"]),
	}
}
