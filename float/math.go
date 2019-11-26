package float

import (
	"github.com/fatih/structs"
	"math"
)

// Add ... Function to add two numbers
func Add(num1, num2 float64) float64 {
	return num1 + num2
}

// Divide ... Function to divide two numbers
func Divide(dividend, divisor float64) float64 {
	return dividend / divisor
}

// Multiply ... Function to multiple two numbers
func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

// Ceil ... Function to return the Ceiling of the number passed.
func Ceil(number float64) float64 {
	return math.Ceil(number)
}

// Floor ... Function to return the Floor of the number passed.
func Floor(number float64) float64 {
	return math.Floor(number)
}

// Max ... Function to return the maximum value in an integer array
func Max(array []float64) float64 {
	var max = -1 * math.MaxFloat64

	for _, elem := range array {
		if elem > max {
			max = elem
		}
	}
	return max
}

// MaxBy ... Function to return the maximum object of the array based on the property specified in the function parameters
func MaxBy(array []interface{}, iteratee string) map[string]interface{} {
	if len(array) <= 0 {
		return nil
	}
	maxVal := -1 * math.MaxFloat64
	var current map[string]interface{}
	for _, elem := range array {
		updateElem := structs.Map(elem)
		floatVal := updateElem[iteratee].(float64)
		if maxVal < floatVal {
			current = updateElem
			maxVal = floatVal
		}
	}
	return current
}

// Min ... Function to return the minimum value in an integer array
func Min(array []float64) float64 {
	var min = math.MaxFloat64

	for _, elem := range array {
		if elem < min {
			min = elem
		}
	}
	return min
}

// MinBy ... Function to return the minimum object of the array based on the property specified in the function parameters
func MinBy(array []interface{}, iteratee string) map[string]interface{} {
	if len(array) <= 0 {
		return nil
	}
	minVal := math.MaxFloat64
	var current map[string]interface{}
	for _, elem := range array {
		updateElem := structs.Map(elem)
		floatVal := updateElem[iteratee].(float64)
		if minVal > floatVal {
			current = updateElem
			minVal = floatVal
		}
	}
	return current
}

// Mean ... Function to return the mean of an array
func Mean(array []float64) float64 {
	return baseSum(array) / 100
}

// MeanBy ... Function to return the mean of an array by it's property value
func MeanBy(array []interface{}, iteratee string) float64 {
	if len(array) <= 0 {
		return 0
	}
	var sum float64
	for _, elem := range array {
		updateElem := structs.Map(elem)
		sum += updateElem[iteratee].(float64)
	}
	return sum
}

// baseSum ... Function to return the sum of all the elements of an array
func baseSum(array []float64) float64 {
	var Sum float64 = 0
	for _, elem := range array {
		Sum += elem
	}
	return Sum
}
