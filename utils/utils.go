package utils

import "errors"

func SplitSlice(numbers []int64, batchSize int64) ([][]int64, error) {
	var begin, end int64

	if batchSize < 0 || len(numbers) == 0 || numbers == nil {
		return nil, errors.New("invalid batch size or empty slice")
	}

	quantity := len(numbers) / int(batchSize)
	if len(numbers)%int(batchSize) != 0 {
		quantity += 1
	}
	end = batchSize
	batches := make([][]int64, quantity)

	for i := 0; i < quantity; i++ {
		if end > int64(len(numbers)) {
			batches[i] = numbers[begin:]
			break
		}
		batches[i] = numbers[begin:end]
		begin = end
		end += batchSize
	}
	return batches, nil
}

func SwapKeyValue(data map[int64]string) (map[string]int64, error) {
	if len(data) == 0 || data == nil {
		return nil, errors.New("map is empty")
	}

	res := make(map[string]int64, len(data))
	for key, value := range data {
		res[value] = key
	}
	return res, nil
}

func SliceFilter(data []int64, filter []int64) ([]int64, error) {
	res := []int64{}
	if filter == nil || data == nil {
		return nil, errors.New("input params invalid")
	}

	filterMap, err := SliceToMap(filter)
	if err != nil {
		return nil, errors.New("failed to convert")
	}

	for i := 0; i < len(data); i++ {
		if _, found := filterMap[data[i]]; !found {
			res = append(res, data[i])
		}
	}
	return res, nil
}

func SliceToMap(data []int64) (map[int64]struct{}, error) {
	if data == nil {
		return nil, errors.New("input params invalid")
	}

	res := make(map[int64]struct{}, len(data))

	for _, val := range data {
		res[val] = struct{}{}
	}
	return res, nil
}
