package processnumber

import "errors"

func ProcessNumber(numbers []int) ([]int, error) {
	// a. Jika input nil, return error
	if numbers == nil {
		return nil, errors.New("no data provided")
	}

	// b. jika daftar kosong, maka panic
	if len(numbers) == 0 {
		panic("empty list provided")
	}

	// c. proses perkalian masing-masing bilangan dengan 2
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num *2
	}

	return result, nil
}