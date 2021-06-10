//package bcalc - basic functions for calculate two values
package bcalc

import (
	"errors"
	"math"
)
//Add - adding two values provided by user, returns result as float64 and error
func Add(a float64, b float64) (float64, error) {
	result := a + b
	return result, nil
}
//Subtract - subtracting two values provided by user, returns result as float64 and error
func Subtract(a float64, b float64) (float64, error) {
	result := a - b
	return result, nil
}
//Multiplication - multiply two values provided by user, returns result as float64 and error
func Multiplication(a float64, b float64) (float64, error) {
	result := a * b
	return result, nil
}
//Divide - divide two values provided by user, returns result as float64 and error
func Divide(a float64, b float64) (float64, error) {
	result := a / b
	return result, nil
}
//Square - square value provided by user, returns result as float64 and error
func Square(a float64) (float64, error) {
	result := a * a
	return result, nil
}
//SquareRoot - root value provided by user. Need to handle error
/*
result, err := bcalc.SquareRoot(-4.2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
returns result as float64 and error
 */
func SquareRoot(a float64) (float64, error) {
	result := math.Sqrt(a)
	if a < 0 {
		return 0, errors.New("Cannot root from negative numbers")
	}
	return result, nil
}
