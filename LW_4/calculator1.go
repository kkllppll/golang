package main

import (
	"fmt"
	"math"
	"net/url"
)

type Task1Calculations struct{}

func (tc *Task1Calculations) Calculate(Sm, Unom, Jek, Ik, Tf, Ct float64) map[string]float64 {
	Im := (Sm / 2) / (math.Sqrt(3) * Unom)
	Ipa := 2 * Im
	Sek := Im / Jek
	Smin := (Ik * math.Sqrt(Tf)) / Ct

	return map[string]float64{
		"Im":   Im,
		"Ipa":  Ipa,
		"Sek":  Sek,
		"Smin": Smin,
	}
}

func calculateTask1(values url.Values) map[string]interface{} {
	sm := parseFloat(values.Get("smController"))
	unom := parseFloat(values.Get("unomController"))
	jek := parseFloat(values.Get("jekController"))
	ik := parseFloat(values.Get("ikController"))
	tf := parseFloat(values.Get("tfController"))
	ct := parseFloat(values.Get("ctController"))

	calculator := Task1Calculations{}
	results := calculator.Calculate(sm, unom, jek, ik, tf, ct)

	return map[string]interface{}{
		"Im":   fmt.Sprintf("%.2f", results["Im"]),
		"Ipa":  fmt.Sprintf("%.2f", results["Ipa"]),
		"Sek":  fmt.Sprintf("%.2f", results["Sek"]),
		"Smin": fmt.Sprintf("%.2f", results["Smin"]),
	}
}
