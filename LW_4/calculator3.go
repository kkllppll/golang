package main

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

type Task3Calculations struct{}

func (tc *Task3Calculations) Calculate(Ukmax, Uvn, Snom, Rcn, Xcn, Rcmin, Xcmin, Unn float64, lengths []float64, rLine, xLine float64) map[string]float64 {
	// реактивний опір трансформатора
	Xt := (Ukmax * Uvn * Uvn) / (100 * Snom)

	// сумарний опір у нормальному режимі
	XsumNorm := Xcn + Xt
	ZshNorm := math.Sqrt(Rcn*Rcn + XsumNorm*XsumNorm)

	// сумарний опір у мінімальному режимі
	XsumMin := Xcmin + Xt
	ZshMin := math.Sqrt(Rcmin*Rcmin + XsumMin*XsumMin)

	// струми трифазного КЗ у нормальному режимі
	Ish3Norm := (Uvn * 1000) / (math.Sqrt(3.0) * ZshNorm)

	// струми двофазного КЗ у нормальному режимі
	Ish2Norm := Ish3Norm * (math.Sqrt(3.0) / 2)

	// струми трифазного КЗ у мінімальному режимі
	Ish3Min := (Uvn * 1000) / (math.Sqrt(3.0) * ZshMin)

	// струми двофазного КЗ у мінімальному режимі
	Ish2Min := Ish3Min * (math.Sqrt(3.0) / 2)

	// коефіцієнт приведення
	kpr := Unn * Unn / (Uvn * Uvn)

	// опори на шинах у нормальному режимі з коефіцієнтом приведення
	RshNorm := Rcn * kpr
	XshNorm := XsumNorm * kpr
	ZshNormPr := math.Sqrt(RshNorm*RshNorm + XshNorm*XshNorm)

	// опори на шинах у мінімальному режимі з коефіцієнтом приведення
	RshMin := Rcmin * kpr
	XshMin := XsumMin * kpr
	ZshMinPr := math.Sqrt(RshMin*RshMin + XshMin*XshMin)

	// дійсні струми трифазного КЗ на шинах 10 кВ у нормальному режимі
	Ish3RealNorm := (Uvn * 1000) / (math.Sqrt(3.0) * ZshNormPr)

	// дійсні струми двофазного КЗ на шинах 10 кВ у нормальному режимі
	Ish2RealNorm := Ish3RealNorm * (math.Sqrt(3.0) / 2)

	// дійсні струми трифазного КЗ на шинах 10 кВ у мінімальному режимі
	Ish3RealMin := (Uvn * 1000) / (math.Sqrt(3.0) * ZshMinPr)

	// дійсні струми двофазного КЗ на шинах 10 кВ у мінімальному режимі
	Ish2RealMin := Ish3RealMin * (math.Sqrt(3.0) / 2)

	// розрахунок сумарних опорів ліній
	RlineSum := 0.0
	XlineSum := 0.0
	for _, length := range lengths {
		RlineSum += rLine * length
		XlineSum += xLine * length
	}

	// сумарний опір у нормальному режимі
	RtotalNorm := RshNorm + RlineSum
	XtotalNorm := XshNorm + XlineSum
	ZtotalNorm := math.Sqrt(RtotalNorm*RtotalNorm + XtotalNorm*XtotalNorm)

	// сумарний опір у мінімальному режимі
	RtotalMin := RshMin + RlineSum
	XtotalMin := XshMin + XlineSum
	ZtotalMin := math.Sqrt(RtotalMin*RtotalMin + XtotalMin*XtotalMin)

	// струми трифазного КЗ
	I3phNorm := (Unn * 1000) / (math.Sqrt(3.0) * ZtotalNorm)
	I3phMin := (Unn * 1000) / (math.Sqrt(3.0) * ZtotalMin)

	// струми двофазного КЗ
	I2phNorm := I3phNorm * (math.Sqrt(3.0) / 2)
	I2phMin := I3phMin * (math.Sqrt(3.0) / 2)

	// підготовка результату
	return map[string]float64{
		"Xt":           Xt,
		"ZshNorm":      ZshNorm,
		"Ish3Norm":     Ish3Norm,
		"Ish2Norm":     Ish2Norm,
		"Ish3Min":      Ish3Min,
		"Ish2Min":      Ish2Min,
		"kpr":          kpr,
		"RshNorm":      RshNorm,
		"XshNorm":      XshNorm,
		"ZshNormPr":    ZshNormPr,
		"RshMin":       RshMin,
		"XshMin":       XshMin,
		"ZshMinPr":     ZshMinPr,
		"Ish3RealNorm": Ish3RealNorm,
		"Ish2RealNorm": Ish2RealNorm,
		"Ish3RealMin":  Ish3RealMin,
		"Ish2RealMin":  Ish2RealMin,
		"RlineSum":     RlineSum,
		"XlineSum":     XlineSum,
		"RtotalNorm":   RtotalNorm,
		"XtotalNorm":   XtotalNorm,
		"ZtotalNorm":   ZtotalNorm,
		"RtotalMin":    RtotalMin,
		"XtotalMin":    XtotalMin,
		"ZtotalMin":    ZtotalMin,
		"I3phNorm":     I3phNorm,
		"I3phMin":      I3phMin,
		"I2phNorm":     I2phNorm,
		"I2phMin":      I2phMin,
		"ZshMin":       ZshMin,
	}
}

func calculateTask3(values url.Values) map[string]interface{} {

	ukmax := parseFloat(values.Get("ukmaxController"))
	uvn := parseFloat(values.Get("uvnController"))
	snom := parseFloat(values.Get("snomController"))
	rcn := parseFloat(values.Get("rcnController"))
	xcn := parseFloat(values.Get("xcnController"))
	rcmin := parseFloat(values.Get("rcminController"))
	xcmin := parseFloat(values.Get("xcminController"))
	unn := parseFloat(values.Get("unnController"))
	lengths := parseCSV(values.Get("lengthsController"))
	rLine := parseFloat(values.Get("rLineController"))
	xLine := parseFloat(values.Get("xLineController"))

	calculator := Task3Calculations{}
	results := calculator.Calculate(ukmax, uvn, snom, rcn, xcn, rcmin, xcmin, unn, lengths, rLine, xLine)

	return map[string]interface{}{
		"Xt":           fmt.Sprintf("%.2f", results["Xt"]),
		"ZshNorm":      fmt.Sprintf("%.2f", results["ZshNorm"]),
		"Ish3Norm":     fmt.Sprintf("%.2f", results["Ish3Norm"]),
		"Ish2Norm":     fmt.Sprintf("%.2f", results["Ish2Norm"]),
		"Ish3Min":      fmt.Sprintf("%.2f", results["Ish3Min"]),
		"Ish2Min":      fmt.Sprintf("%.2f", results["Ish2Min"]),
		"kpr":          fmt.Sprintf("%.4f", results["kpr"]),
		"RshNorm":      fmt.Sprintf("%.2f", results["RshNorm"]),
		"XshNorm":      fmt.Sprintf("%.2f", results["XshNorm"]),
		"ZshNormPr":    fmt.Sprintf("%.2f", results["ZshNormPr"]),
		"RshMin":       fmt.Sprintf("%.2f", results["RshMin"]),
		"XshMin":       fmt.Sprintf("%.2f", results["XshMin"]),
		"ZshMinPr":     fmt.Sprintf("%.2f", results["ZshMinPr"]),
		"Ish3RealNorm": fmt.Sprintf("%.2f", results["Ish3RealNorm"]),
		"Ish2RealNorm": fmt.Sprintf("%.2f", results["Ish2RealNorm"]),
		"Ish3RealMin":  fmt.Sprintf("%.2f", results["Ish3RealMin"]),
		"Ish2RealMin":  fmt.Sprintf("%.2f", results["Ish2RealMin"]),
		"RlineSum":     fmt.Sprintf("%.2f", results["RlineSum"]),
		"XlineSum":     fmt.Sprintf("%.2f", results["XlineSum"]),
		"RtotalNorm":   fmt.Sprintf("%.2f", results["RtotalNorm"]),
		"XtotalNorm":   fmt.Sprintf("%.2f", results["XtotalNorm"]),
		"ZtotalNorm":   fmt.Sprintf("%.2f", results["ZtotalNorm"]),
		"RtotalMin":    fmt.Sprintf("%.2f", results["RtotalMin"]),
		"XtotalMin":    fmt.Sprintf("%.2f", results["XtotalMin"]),
		"ZtotalMin":    fmt.Sprintf("%.2f", results["ZtotalMin"]),
		"I3phNorm":     fmt.Sprintf("%.2f", results["I3phNorm"]),
		"I3phMin":      fmt.Sprintf("%.2f", results["I3phMin"]),
		"I2phNorm":     fmt.Sprintf("%.2f", results["I2phNorm"]),
		"I2phMin":      fmt.Sprintf("%.2f", results["I2phMin"]),
		"ZshMin":       fmt.Sprintf("%.2f", results["ZshMin"]),
	}
}

func parseFloat(value string) float64 {
	value = strings.Replace(value, ",", ".", -1)
	num, _ := strconv.ParseFloat(value, 64)
	return num
}

func parseCSV(value string) []float64 {
	parts := strings.Split(value, ",")
	var result []float64
	for _, p := range parts {
		result = append(result, parseFloat(strings.TrimSpace(p)))
	}
	return result
}
