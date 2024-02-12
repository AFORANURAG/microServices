package utils

import (
	"reflect"
)

func ExtractKeysAndTypes(data interface{}) map[string]reflect.Type {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Struct {
		panic("Input is not a struct")
	}

	numFields := value.NumField()
	keysAndTypes := make(map[string]reflect.Type)

	for i := 0; i < numFields; i++ {
		field := value.Type().Field(i)
		keysAndTypes[field.Name] = field.Type
	}

	return keysAndTypes
}
