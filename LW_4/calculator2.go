package main

import (
	"fmt"
	"math"
	"net/url"
)

type Task2Calculations struct{}

func (tc *Task2Calculations) Calculate(Usn, Sk, Uk, Snom float64) map[string]float64 {
	Xc := math.Pow(Usn, 2) / Sk
	Xt := Uk / 100 * math.Pow(Usn, 2) / Snom
	Xsum := Xc + Xt
	Ip0 := Usn / (math.Sqrt(3) * Xsum)

	return map[string]float64{
		"Xc":   Xc,
		"Xt":   Xt,
		"Xsum": Xsum,
		"Ip0":  Ip0,
	}
}

func calculateTask2(values url.Values) map[string]interface{} {
	usn := parseFloat(values.Get("usnController"))
	sk := parseFloat(values.Get("skController"))
	uk := parseFloat(values.Get("ukController"))
	snom := parseFloat(values.Get("snomController"))

	calculator := Task2Calculations{}
	results := calculator.Calculate(usn, sk, uk, snom)

	return map[string]interface{}{
		"Xc":   fmt.Sprintf("%.2f", results["Xc"]),
		"Xt":   fmt.Sprintf("%.2f", results["Xt"]),
		"Xsum": fmt.Sprintf("%.2f", results["Xsum"]),
		"Ip0":  fmt.Sprintf("%.2f", results["Ip0"]),
	}
}
