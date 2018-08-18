package main

import (
	"fmt"
	"reflect"
	"strings"
)

func ConcatStr(params []string) string {
	var result string
	for _, param := range params {
		result += param
	}
	return result
}

func ConcatStr1(params ...interface{}) string {
	var params1 []string
	paramsType := reflect.TypeOf(params)
	if paramsType.Kind() == reflect.Slice || paramsType.Kind() == reflect.Array {
		params1 = make([]string, len(params))
		for _, param := range params {
			params1 = append(params1, param.(string))
		}
	}

	return strings.Join(params1, "")
}

func main() {
	fmt.Println("params: 1,2,3")
	result := ConcatStr([]string{"1", "2", "3"})
	fmt.Printf("result: %s", result)
}
