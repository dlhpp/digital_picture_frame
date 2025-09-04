package utils

import (
	"fmt"
	"reflect"

	"github.com/dlhpp/digital_picture_frame/logging"
)

func DescribeVariable(name string, i any) string {
	if i == nil {
		return "nil"
	}

	r := reflect.TypeOf(i)

	return fmt.Sprintf("(name=%s, reflectName=%s, kind=%s, type=%T, value=%+v)", name, r.Name(), r.Kind(), i, i)
}

func ConvertToSliceStrings(interfaceSlice []interface{}) []string {
	// Our slice of interfaces is expected to contain strings
	// TODO: I should check that reflect.TypeOf(interfaceSlice).Kind() == reflect.Slice first

	stringSlice := make([]string, len(interfaceSlice))

	// Iterate through the interface slice and perform type assertion on each element
	for i, v := range interfaceSlice {
		s, ok := v.(string)
		if !ok {
			// Handle the case where an element is not a string (e.g., panic, return error)
			panic(fmt.Sprintf("ConvertToSliceStrings: Not a string at i=%d, %s \n", i, DescribeVariable("v", v)))
		}
		stringSlice[i] = s
	}

	logging.Log("ConvertToSliceStrings", 2, DescribeVariable("stringSlice", stringSlice))

	return stringSlice
}

func ConvertToSliceInt(interfaceSlice []interface{}) []int {
	// Our slice of interfaces is expected to contain int
	// TODO: I should check that reflect.TypeOf(interfaceSlice).Kind() == reflect.Slice first

	intSlice := make([]int, len(interfaceSlice))

	// Iterate through the interface slice and perform type assertion on each element
	for i, v := range interfaceSlice {
		s, ok := v.(int)
		if !ok {
			// Handle the case where an element is not a string (e.g., panic, return error)
			panic(fmt.Sprintf("ConvertToSliceInt: Not an int at i=%d, %s \n", i, DescribeVariable("v", v)))
		}
		intSlice[i] = s
	}

	logging.Log("ConvertToSliceInt", 2, DescribeVariable("intSlice", intSlice))

	return intSlice
}

func ConvertToSliceFloat(interfaceSlice []interface{}) []float64 {
	// Our slice of interfaces is expected to contain int
	// TODO: I should check that reflect.TypeOf(interfaceSlice).Kind() == reflect.Slice first

	floatSlice := make([]float64, len(interfaceSlice))

	// Iterate through the interface slice and perform type assertion on each element
	for i, v := range interfaceSlice {
		s, ok := v.(float64)
		if !ok {
			// Handle the case where an element is not a string (e.g., panic, return error)
			panic(fmt.Sprintf("ConvertToSliceFloat: Not a float64 at i=%d, %s \n", i, DescribeVariable("v", v)))
		}
		floatSlice[i] = s
	}

	logging.Log("ConvertToSliceFloat", 2, DescribeVariable("floatSlice", floatSlice))

	return floatSlice
}

// func ConvertToMap(interfaceSlice interface{}) []float64 {
// 	// Our slice of interfaces is expected to contain int
// 	// TODO: I should check that reflect.TypeOf(interfaceSlice).Kind() == reflect.Slice first

// 	floatSlice := make([]float64, len(interfaceSlice))

// 	// Iterate through the interface slice and perform type assertion on each element
// 	for i, v := range interfaceSlice {
// 		s, ok := v.(float64)
// 		if !ok {
// 			// Handle the case where an element is not a string (e.g., panic, return error)
// 			panic(fmt.Sprintf("ConvertToSliceFloat: Not a float64 at i=%d, %s \n", i, DescribeVariable("v", v)))
// 		}
// 		floatSlice[i] = s
// 	}

// 	logging.Log("ConvertToSliceFloat", 2, DescribeVariable("floatSlice", floatSlice))

// 	return floatSlice
// }
