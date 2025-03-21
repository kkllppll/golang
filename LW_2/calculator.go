package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// структура для розрахунку викидів
type FuelCalculator struct {
	Q_iR float64 // нижча теплота згоряння (МДж/кг)
	L    float64 // частка леткої золи
	Ar   float64 // масова частка золи Ar (%)
	G    float64 // гвин (%)
	n    float64 // ефективність золовловлення ηзу
	k    float64 // Ктв
	B    float64 // витрати палива B (т)
}

// отримання вхідних даних з форми
func (fc *FuelCalculator) InputParameters(values url.Values) {
	fc.Q_iR = parseFloat(values.Get("Q_iR"))
	fc.L = parseFloat(values.Get("L"))
	fc.Ar = parseFloat(values.Get("Ar"))
	fc.G = parseFloat(values.Get("G"))
	fc.n = parseFloat(values.Get("n"))
	fc.k = parseFloat(values.Get("k"))
	fc.B = parseFloat(values.Get("B"))
}

// розрахунок показника емісії (г/ГДж)
func (fc *FuelCalculator) CalculateEmissionFactor() (float64, error) {
	if fc.Q_iR == 0 {
		return 0, fmt.Errorf("Нижча теплота згоряння не може бути нульовою")
	}
	return (1_000_000/fc.Q_iR)*fc.L*(fc.Ar/(100-fc.G))*(1-fc.n) + fc.k, nil
}

// розрахунок валового викиду (т)
func (fc *FuelCalculator) CalculateTotalEmission() (float64, error) {
	emissionFactor, err := fc.CalculateEmissionFactor()
	if err != nil {
		return 0, err
	}
	return 0.000001 * emissionFactor * fc.Q_iR * fc.B, nil
}

// calculateTask – виконує всі розрахунки для задачі
func calculateTask(values url.Values) map[string]interface{} {
	calculator := FuelCalculator{}
	calculator.InputParameters(values)

	emissionFactor, err1 := calculator.CalculateEmissionFactor()
	totalEmission, err2 := calculator.CalculateTotalEmission()

	// Отримуємо тип палива
	fuelType := values.Get("fuelType")

	if err1 != nil || err2 != nil {
		return map[string]interface{}{
			"error": "Помилка в розрахунках: " + err1.Error(),
		}
	}

	return map[string]interface{}{
		"fuelType":       fuelType,
		"emissionFactor": fmt.Sprintf("%.2f г/ГДж", emissionFactor),
		"totalEmission":  fmt.Sprintf("%.2f т", totalEmission),
	}
}

// конвертує текст у число
func parseFloat(value string) float64 {
	value = strings.Replace(value, ",", ".", -1)
	num, _ := strconv.ParseFloat(value, 64)
	return num
}
