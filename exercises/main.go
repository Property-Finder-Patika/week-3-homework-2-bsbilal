package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Rectangle struct {
	aSide         float64
	bSide         float64
	area          float64
	circumference float64
}

func main() {

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please provide 2 values.")
		return
	}

	err, newRectangle := CreateRectangle(args[0:2])
	if err != nil {
		fmt.Println(err)
	}

	SetRectangleArea(&newRectangle)

}
func CreateRectangle(arr []string) (error, Rectangle) {
	var numberArr [2]float64
	for i, _ := range arr {
		tmp, err := strconv.ParseFloat(arr[i], 32)
		if err != nil {
			return errors.New("The %d number as expected. It should be positive decimal number."), Rectangle{}
		} else {
			if i < 0 {
				return errors.New("The %d number is should be a positive."), Rectangle{}
			} else {
				numberArr[i] = tmp
			}
		}
	}

	r := Rectangle{aSide: numberArr[0], bSide: numberArr[1]}

	return nil, r
}

func SetRectangleArea(r *Rectangle) {
	r.area = r.aSide * r.bSide
	return
}
