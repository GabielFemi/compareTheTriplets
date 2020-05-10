package main

import (
	"bufio"
	"fmt"
	"os"
	_"reflect"
	"strconv"
	"strings"
)

func main() {
	aliScanner := bufio.NewScanner(os.Stdin)
	aliScanner.Scan()
	aliScannedValue := aliScanner.Text()
	aliRating := strings.Fields(aliScannedValue)
	fmt.Println("Ali's rating:", aliRating)

	var aliConvertedRating []int
	// Get the rating and convert it to an integer
	for _, i:= range aliRating {
		convertToInteger, err := strconv.Atoi(i)
		aliConvertedRating = append(aliConvertedRating, convertToInteger)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(i, "type:", reflect.TypeOf(convertToInteger))
	}
	//
}

func compareTriplets() {

}