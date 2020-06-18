package format

import (
	"errors"
	"fmt"
	"strings"
)

type Items map[string]interface{}

func String(template string, items Items) string {
	for key, value := range items {
		template = strings.ReplaceAll(template, fmt.Sprintf("{%v}", key), fmt.Sprintf("%v", value))
	}
	return template
}

func Error(template string, items Items) error {
	template = String(template, items)
	return errors.New(template)
}

func Println(template string, items Items) {
	template = String(template, items)
	fmt.Println(template)
}
