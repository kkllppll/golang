package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type FuelCalculator struct {
	HydrogenP float64
	CarbonP   float64
	SulfurP   float64
	NitrogenP float64
	OxygenP   float64
	WaterP    float64
	AshP      float64
}

func (fc *FuelCalculator) InputFuelComponents(values url.Values) {
	fc.HydrogenP = parseFloat(values.Get("hydrogen"))
	fc.CarbonP = parseFloat(values.Get("carbon"))
	fc.SulfurP = parseFloat(values.Get("sulfur"))
	fc.NitrogenP = parseFloat(values.Get("nitrogen"))
	fc.OxygenP = parseFloat(values.Get("oxygen"))
	fc.WaterP = parseFloat(values.Get("water"))
	fc.AshP = parseFloat(values.Get("ash"))
}

func (fc *FuelCalculator) CoeffWorkingToDry() (map[string]float64, float64) {
	coeff := 100 / (100 - fc.WaterP)
	composition := map[string]float64{
		"Hydrogen": fc.HydrogenP * coeff,
		"Carbon":   fc.CarbonP * coeff,
		"Sulfur":   fc.SulfurP * coeff,
		"Nitrogen": fc.NitrogenP * coeff,
		"Oxygen":   fc.OxygenP * coeff,
		"Ash":      fc.AshP * coeff,
	}
	return composition, coeff
}

func (fc *FuelCalculator) coeffWorkingToComb() (map[string]float64, float64) {
	coeff := 100 / (100 - fc.WaterP - fc.AshP)
	composition := map[string]float64{
		"Hydrogen": fc.HydrogenP * coeff,
		"Carbon":   fc.CarbonP * coeff,
		"Sulfur":   fc.SulfurP * coeff,
		"Nitrogen": fc.NitrogenP * coeff,
		"Oxygen":   fc.OxygenP * coeff,
	}
	return composition, coeff
}

func (fc *FuelCalculator) CalculateLowerHeatingValueWorkingMass() float64 {
	return 339*fc.CarbonP + 1030*fc.HydrogenP - 108.8*(fc.OxygenP-fc.SulfurP) - 25*fc.WaterP
}

func (fc *FuelCalculator) calculateLowerHeatingValueDryMass(workingQ float64) float64 {
	return (workingQ + 0.025*fc.WaterP) * 100 / (100 - fc.WaterP)
}

func (fc *FuelCalculator) calculateLowerHeatingValueCombustibleMass(workingQ float64) float64 {
	return (workingQ + 0.025*fc.WaterP) * 100 / (100 - fc.WaterP - fc.AshP)
}

func calculateTask1(values url.Values) map[string]interface{} {
	calculator := FuelCalculator{}
	calculator.InputFuelComponents(values)

	dryComposition, dryCoeff := calculator.CoeffWorkingToDry()
	combComposition, combCoeff := calculator.coeffWorkingToComb()
	workingQ := calculator.CalculateLowerHeatingValueWorkingMass()
	dryQ := calculator.calculateLowerHeatingValueDryMass(workingQ)
	combQ := calculator.calculateLowerHeatingValueCombustibleMass(workingQ)

	return map[string]interface{}{
		"coefficient_dry":  fmt.Sprintf("%.2f", dryCoeff),
		"composition_dry":  formatComposition(dryComposition),
		"coefficient_comb": fmt.Sprintf("%.2f", combCoeff),
		"composition_comb": formatComposition(combComposition),
		"workingQ":         fmt.Sprintf("%.1f", workingQ),
		"dryQ":             fmt.Sprintf("%.1f", dryQ),
		"combQ":            fmt.Sprintf("%.1f", combQ),
	}
}

func parseFloat(value string) float64 {
	value = strings.Replace(value, ",", ".", -1)
	num, _ := strconv.ParseFloat(value, 64)
	return num
}

func formatComposition(comp map[string]float64) string {
	var result strings.Builder
	for key, value := range comp {
		result.WriteString(fmt.Sprintf("%s: %.2f\n", key, value))
	}
	return result.String()
}
