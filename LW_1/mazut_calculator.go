package main

import (
	"fmt"
	"net/url"
)

type MazutCalculator struct {
	HydrogenP    float64
	CarbonP      float64
	SulfurP      float64
	OxygenP      float64
	VanadiumP    float64
	WaterP       float64
	AshP         float64
	LowerHeating float64
}

func (mc *MazutCalculator) InputFuelComponents(values url.Values) {
	mc.HydrogenP = parseFloat(values.Get("hydrogen"))
	mc.CarbonP = parseFloat(values.Get("carbon"))
	mc.SulfurP = parseFloat(values.Get("sulfur"))
	mc.OxygenP = parseFloat(values.Get("oxygen"))
	mc.VanadiumP = parseFloat(values.Get("vanadium"))
	mc.WaterP = parseFloat(values.Get("water"))
	mc.AshP = parseFloat(values.Get("ash"))
	mc.LowerHeating = parseFloat(values.Get("lowerHeating"))
}

func (mc *MazutCalculator) CombToWorking() map[string]float64 {
	coeff := (100 - mc.WaterP - mc.AshP) / 100
	return map[string]float64{
		"Carbon":   mc.CarbonP * coeff,
		"Hydrogen": mc.HydrogenP * coeff,
		"Oxygen":   mc.OxygenP * coeff,
		"Sulfur":   mc.SulfurP * coeff,
		"Ash":      mc.AshP * (100 - mc.WaterP) / 100,
		"Vanadium": mc.VanadiumP * (100 - mc.WaterP) / 100,
	}
}

func (mc *MazutCalculator) CalculateWorkingHeatingValue() float64 {
	return mc.LowerHeating*(100-mc.WaterP-mc.AshP)/100 - 0.025*mc.WaterP
}

func calculateTask2(values url.Values) map[string]interface{} {
	calculator := MazutCalculator{}
	calculator.InputFuelComponents(values)

	composition := calculator.CombToWorking()
	lowerHeating := calculator.CalculateWorkingHeatingValue()

	return map[string]interface{}{
		"composition":  formatComposition(composition),
		"lowerHeating": fmt.Sprintf("%.2f", lowerHeating),
	}
}
