package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
)

const myURL = "http://127.0.0.1:7777"

func main() {
	reader := bufio.NewReader(os.Stdin)
	var variables [4]string
	for i := 0; i < len(variables); i++ {
		input, _ := reader.ReadString('\n')
		variables[i] = strings.TrimSpace(input)
	}

	allowMethods("MEW")

	yahoo, _ := http.Get(myURL)
	listOfValues := make([][]interface{}, 0)
	for i := 0; i < 3; i++ {
		count := 3
		var variable string

		switch i {
		case 2:
			variable = variables[i]
			count = 1
		default:
			variable = variables[i] + "," + variables[i+1] + "," + variables[i+2]
		}
		con := yahoo
		con.Header.Set("X-cat-VAriable", variable)
		con.Method = "MEW"
		response, _ := con.Do()
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			headers := response.Header
			var headerValues []string
			for key := range headers {
				if key == "" {
					continue
				}
				if strings.ToUpper(key) == "X-CAT-VALUE" {
					headerValues = headers[key]
					break
				}
			}
			b := make([]interface{}, count)
			for _, value := range headerValues {
				b = append(b, value)
			}
			listOfValues = append(listOfValues, b)
		}
	}
	values := make([]string, 4)
	first := make([]string, len(listOfValues[0]))
	copy(first, listOfValues[0]) // 1 2 3
	second := make([]string, len(listOfValues[1]))
	copy(second, listOfValues[1]) // 2 3 4
	third := listOfValues[2]

	values[2] = fmt.Sprintf("%v", third[0])
	pair := make([]string, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if first[i] == second[j] {
				pair = append(pair, first[i])
				first[i] = ""
				second[j] = ""
				break
			}
		}
		if len(pair) == 2 {
			break
		}
	}

	first = make([]string, len(listOfValues[0]))
	copy(first, listOfValues[0])
	removeDuplicates(first, pair)
	values[0] = first[0]

	second = make([]string, len(listOfValues[1]))
	copy(second, listOfValues[1])
	removeDuplicates(second, pair)
	values[3] = second[0]

	a := make([]string, 1)
	a[0] = values[2]
	removeDuplicates(pair, a)
	values[1] = pair[0]

	for _, v := range values {
		fmt.Println(v)
	}
}

func removeDuplicates(first []string, second []string) {
	for _, value := range second {
		for i := 0; i < len(first); i++ {
			if first[i] == value {
				first = append(first[:i], first[i+1:]...)
				break
			}
		}
	}
}

func allowMethods(meth string) {
	methodsField := reflect.ValueOf(http.DefaultClient).Elem().FieldByName("methods")
	methodsField.Set(reflect.AppendSlice(methodsField, reflect.ValueOf([]string{meth})))
}
