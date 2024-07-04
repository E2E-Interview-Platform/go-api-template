package helpers

import (
	"errors"
	"sort"
)

func GetIndexOfElementInSlice(elements []int, item int) (int, error) {
	for i := range elements {
		if elements[i] == item {
			return i, nil
		}
	}
	return -1, errors.New("element not found")
}

func SortIntDescending(elements []int) {
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] > elements[j]
	})
}
