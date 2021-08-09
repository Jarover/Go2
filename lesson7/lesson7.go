package lesson7

import (
	"fmt"
	"reflect"
	"strings"
)

//InQuery - получет строку sql запроса и преобразует ее, в зависимости от аргументов
func InQuery(query string, args ...interface{}) (string, []interface{}) {

	var outArgs []interface{}

	outQuery := ""
	// разбиваем строку на подстроки по символу вопроса
	arr := strings.Split(query, "?")

	// цикл по агрументам
	for i, n := range args {

		p := parseArg(n)
		outQuery = outQuery + arr[i]

		for j := 0; j < len(p); j++ {
			outArgs = append(outArgs, p[j])
			if j > 0 {
				outQuery = outQuery + ",?"
			} else {
				outQuery = outQuery + "?"
			}

		}

	}
	fmt.Println(outQuery)
	fmt.Println(outArgs)
	return outQuery, outArgs
}

//parseArg - принимает аргумент и возвращает слайс интерфесов
func parseArg(arg interface{}) []interface{} {
	var out []interface{}
	valueOf := reflect.ValueOf(arg)
	fmt.Println("value:", valueOf.Kind())

	// определяем тип
	switch v := reflect.ValueOf(arg); v.Kind() {
	case reflect.Slice:

		s := reflect.ValueOf(arg)

		for i := 0; i < s.Len(); i++ {

			out = append(out, s.Index(i))
		}
	case reflect.Bool, reflect.String, reflect.Float32, reflect.Float64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out = append(out, arg)

	default:
		fmt.Printf("Неопознаный тип %s", v.Kind())
	}

	return out
}
