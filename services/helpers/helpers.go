package helpers

import (
	"strconv"
	"strings"
)

func StringToNumArray(A string) (result []int, err error) {
	a := strings.Split(A, ",")
	for _, v := range a {
		num, err := strconv.Atoi(v)
		result = append(result, num)
		if err != nil {
			return []int{}, err
		}
	}

	return
}
