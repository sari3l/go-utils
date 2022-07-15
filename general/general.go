package general

import "reflect"

// map | struct

func Keys(src any) []any {
	var keys []any
	switch reflect.ValueOf(src).Type().Kind() {
	case reflect.Map:
		for _, v := range reflect.ValueOf(src).MapKeys() {
			keys = append(keys, v.Interface())
		}
	case reflect.Struct:
		for i := 0; i < reflect.ValueOf(src).Type().NumField(); i++ {
			keys = append(keys, reflect.ValueOf(src).Type().Field(i).Name)
		}
	}
	return keys
}

func Get(src any, dst string, defaultValue any) any {
	index := Index(Keys(src), dst)
	if index > -1 {
		switch reflect.ValueOf(src).Type().Kind() {
		case reflect.Map:
			return reflect.ValueOf(src).MapIndex(reflect.ValueOf(dst)).Interface()
		case reflect.Struct:
			if reflect.TypeOf(src).Field(index).IsExported() {
				return reflect.ValueOf(src).FieldByName(dst).Interface()
			}
		}
	}
	return defaultValue
}

// slice

func Contains[T, E any](src []T, dst E) bool {
	return Index(src, dst) > -1
}

func Index[T, E any](src []T, dst E) int {
	for i, s := range src {
		if reflect.TypeOf(s).Kind() == reflect.TypeOf(dst).Kind() {
			if reflect.ValueOf(s).Interface() == reflect.ValueOf(dst).Interface() {
				return i
			}
		}
	}
	return -1
}
