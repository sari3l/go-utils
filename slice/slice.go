package slice

import "github.com/sari3l/go-utils/general"

func Contains[T comparable](src []T, dst T) bool {
	return general.Index(src, dst) > -1
}

func Index[T comparable](src []T, dst T) int {
	return general.Index(src, dst)
}
