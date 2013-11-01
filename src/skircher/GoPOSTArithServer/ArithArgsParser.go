// ArithArgsParser
package main

import (
	//"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseArgsInt(args string) ([]int, error) {

	//if args != nil {
	//	return nil, errors.New("parse Nil")
	//}

	strSlice := strings.Split(args, ",")
	iSlice := make([]int, len(strSlice))
	fmt.Println("Hello Strings", strSlice, len(strSlice))
	var err error
	var num int = 0
	for i, n := range strSlice {
		//fmt.Println("Hello loop %v %s", i, n)
		num, err = strconv.Atoi(strings.TrimSpace(n))
		fmt.Println("Hello parsed %v %s", num, n)
		if err != nil {
			fmt.Println("Hello error", err)
		} else {
			iSlice[i] = num
		}
	}
	fmt.Println("Hello Ints", iSlice)
	return iSlice, err
}
func parseArgsF64(args string) ([]float64, error) {

	//if args != nil {
	//	return nil, errors.New("parse Nil")
	//}

	strSlice := strings.Split(args, ",")
	iSlice := make([]float64, len(strSlice))
	fmt.Println("Hello Strings", strSlice, len(strSlice))
	var err error
	var num float64
	for i, n := range strSlice {
		//fmt.Println("Hello loop %v %s", i, n)
		num, err = strconv.ParseFloat(strings.TrimSpace(n), 64)
		fmt.Println("Hello parsed %v %s", num, n)
		if err != nil {
			fmt.Println("Hello error", err)
		} else {
			iSlice[i] = num
		}
	}
	fmt.Println("Hello Floats", iSlice)
	return iSlice, err
}
