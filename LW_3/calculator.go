package main

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

// структура для розрахунку викидів
type Calculator struct {
	pc      float64 // нижча теплота згоряння (МДж/кг)
	sigmaP  float64 // частка леткої золи
	sigmaLP float64 // масова частка золи Ar (%)
	b       float64 // гвин (%)
}

// отримання вхідних даних з форми
func (fc *Calculator) InputParameters(values url.Values) {
	fc.pc = parseFloat(values.Get("powerController"))
	fc.sigmaP = parseFloat(values.Get("deviationController"))
	fc.sigmaLP = parseFloat(values.Get("lowerDeviationController"))
	fc.b = parseFloat(values.Get("costController"))
}
func (c *Calculator) pD(p, sigma float64) float64 {
	return (1 / (sigma * math.Sqrt(2*math.Pi))) * math.Exp(-math.Pow(p-c.pc, 2)/(2*math.Pow(sigma, 2)))
}

func (c *Calculator) calculateDeltaW(pMin, pMax, deltaP, sigma float64) float64 {
	sum := 0.0
	p := pMin
	for p <= pMax {
		sum += c.pD(p, sigma) * deltaP
		p += deltaP
	}
	return sum
}

func (c *Calculator) calculateEnergy(deltaW float64) float64 {
	return deltaW * c.pc * 24
}

func (c *Calculator) calculatePenaltyEnergy(deltaW float64) float64 {
	return c.pc * 24 * (1 - deltaW)
}

func (c *Calculator) calculateProfit(w float64) float64 {
	return w * c.b
}

func (c *Calculator) calculatePenalty(wTotal, wSafe float64) float64 {
	return (wTotal - wSafe) * c.b
}

func (c *Calculator) Calculate() map[string]interface{} {
	deltaP := 0.001
	pMin := 4.75
	pMax := 5.25

	deltaW1 := c.calculateDeltaW(pMin, pMax, deltaP, c.sigmaP)
	deltaW2 := c.calculateDeltaW(pMin, pMax, deltaP, c.sigmaLP)

	w1 := c.calculateEnergy(deltaW1)
	w2 := c.calculateEnergy(deltaW2)

	wPenalty1 := c.calculatePenaltyEnergy(deltaW1)
	wPenalty2 := c.calculatePenaltyEnergy(deltaW2)

	profit1 := c.calculateProfit(w1)
	profit2 := c.calculateProfit(w2)

	wTotal := c.calculateEnergy(1.0)
	penalty1 := c.calculatePenalty(wTotal, w1)
	penalty2 := c.calculatePenalty(wTotal, w2)

	netProfit1 := profit1 - penalty1
	netProfit2 := profit2 - penalty2

	return map[string]interface{}{
		"delta_W1":   fmt.Sprintf("%.2f", deltaW1*100),
		"W1":         fmt.Sprintf("%.2f", w1),
		"profit1":    fmt.Sprintf("%.2f", profit1),
		"penalty1":   fmt.Sprintf("%.2f", penalty1),
		"netProfit1": fmt.Sprintf("%.2f", netProfit1),
		"W_penalty1": fmt.Sprintf("%.2f", wPenalty1),
		"delta_W2":   fmt.Sprintf("%.2f", deltaW2*100),
		"W2":         fmt.Sprintf("%.2f", w2),
		"profit2":    fmt.Sprintf("%.2f", profit2),
		"penalty2":   fmt.Sprintf("%.2f", penalty2),
		"netProfit2": fmt.Sprintf("%.2f", netProfit2),
		"W_penalty2": fmt.Sprintf("%.2f", wPenalty2),
	}
}

// parseFloat - конвертує текст у число
func parseFloat(value string) float64 {
	value = strings.Replace(value, ",", ".", -1)
	num, _ := strconv.ParseFloat(value, 64)
	return num
}
