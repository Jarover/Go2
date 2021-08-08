package lesson7

import (
	"fmt"
	"reflect"
	"strings"
)

//InQuery - получет строку sql запроса и преобразует ее, в зависимости от аргументов
func InQuery(query string, args ...interface{}) (string,[]interface{}) {

	var outArgs  []interface{}
	out := ""
	arr := strings.Split(query,"?")

	for i, n := range args {

		p := parseArg(n)
		out = out + arr[i]
		for j := 0; j < len(p); j++ {
			outArgs = append(outArgs,p[j])
			if j>0 {
				out = out + ",?"
			} else {
				out = out + "?"
			}

		}

	}
	fmt.Println(out)
	fmt.Println(outArgs)
	return query,outArgs
}

func parseArg(arg interface{}) []interface{} {
	var out []interface{}
	valueOf := reflect.ValueOf(arg)
	fmt.Println("value:", valueOf.Kind())


	switch v := reflect.ValueOf(arg); v.Kind() {
	case reflect.Slice:

		s := reflect.ValueOf(arg)

		for i := 0; i < s.Len(); i++ {

			out = append(out, s.Index(i))
		}
	case reflect.Bool,  reflect.String,reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out = append(out, arg)

	default:
		fmt.Printf("unhandled kind %s", v.Kind())
	}


	return out
}
