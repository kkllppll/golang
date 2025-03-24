package main

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
)

type Calculations3 struct{}

// getSimultaneityCoefficient обчислює коефіцієнт одночасності
func (c *Calculations3) getSimultaneityCoefficient(Kv float64, connections int) float64 {
	kvRanges := []float64{0.3, 0.5, 0.8}
	koTable := [][]float64{
		{0.9, 0.8, 0.75, 0.7},  // Kv < 0.3
		{0.95, 0.9, 0.85, 0.8}, // 0.3 ≤ Kv < 0.5
		{1.0, 0.95, 0.9, 0.85}, // 0.5 ≤ Kv < 0.8
		{1.0, 1.0, 0.95, 0.9},  // Kv ≥ 0.8
	}

	var kvIndex int
	if Kv < kvRanges[0] {
		kvIndex = 0
	} else if Kv < kvRanges[1] {
		kvIndex = 1
	} else if Kv < kvRanges[2] {
		kvIndex = 2
	} else {
		kvIndex = 3
	}

	var connectionsIndex int
	if connections <= 4 {
		connectionsIndex = 0
	} else if connections <= 8 {
		connectionsIndex = 1
	} else if connections <= 25 {
		connectionsIndex = 2
	} else {
		connectionsIndex = 3
	}

	return koTable[kvIndex][connectionsIndex]
}

func (c *Calculations3) Calculate(quantity int, Phi, nPhKv, nPnKvtg, np2 float64) map[string]float64 {
	Kv := 0.0
	if Phi > 0 {
		Kv = nPhKv / Phi
	}

	Ne := 0.0
	if np2 > 0 {
		Ne = (Phi * Phi) / np2
	}

	Kp := c.getSimultaneityCoefficient(Kv, quantity)
	Pp := Kp * nPhKv
	Qp := Kp * nPnKvtg
	Sp := math.Sqrt(Pp*Pp + Qp*Qp)
	Uh := 0.38
	Ip := Pp / Uh

	return map[string]float64{
		"Kv": Kv,
		"Ne": Ne,
		"Kp": Kp,
		"Pp": Pp,
		"Qp": Qp,
		"Sp": Sp,
		"Ip": Ip,
	}
}

// parseInt допоміжна функція для парсингу цілих чисел із рядка
func parseInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Помилка парсингу:", value)
		return 0 // значення за замовчуванням
	}
	return num
}

// отримує значення   та виконує обчислення
func calculateTask3(values url.Values) map[string]interface{} {
	quantity := parseInt(values.Get("quantityController"))
	Phi := parseFloat(values.Get("phiController"))
	nPhKv := parseFloat(values.Get("sulnPhKvControllerfur"))
	nPnKvtg := parseFloat(values.Get("nPnKvtgController"))
	np2 := parseFloat(values.Get("np2Controller"))

	calculator := Calculations3{}
	results := calculator.Calculate(quantity, Phi, nPhKv, nPnKvtg, np2)

	return map[string]interface{}{
		"Kv": fmt.Sprintf("%.2f", results["Kv"]),
		"Ne": fmt.Sprintf("%.2f", results["Ne"]),
		"Kp": fmt.Sprintf("%.2f", results["Kp"]),
		"Pp": fmt.Sprintf("%.2f", results["Pp"]),
		"Qp": fmt.Sprintf("%.2f", results["Qp"]),
		"Sp": fmt.Sprintf("%.2f", results["Sp"]),
		"Ip": fmt.Sprintf("%.2f", results["Ip"]),
	}
}
