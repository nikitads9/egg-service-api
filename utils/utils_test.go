package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitSlice(t *testing.T) {
	t.Run("last batch not full", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := [][]int64{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}}
		res, err := SplitSlice(numbers, int64(2))
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("last batch not full", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8}
		expected := [][]int64{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
		res, err := SplitSlice(numbers, int64(2))
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("batch size is equal to slice size", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8}
		expected := [][]int64{{1, 2, 3, 4, 5, 6, 7, 8}}
		res, err := SplitSlice(numbers, int64(8))
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty slice passed", func(t *testing.T) {
		numbers := []int64{}
		_, err := SplitSlice(numbers, int64(2))
		require.Error(t, err)
	})

	t.Run("negative batchSize passed", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
		_, err := SplitSlice(numbers, int64(-2))
		require.Error(t, err)
	})
}

func TestSwapKeyValue(t *testing.T) {
	t.Run("swap successful", func(t *testing.T) {
		data := map[int64]string{1: "адын", 2: "цвай", 3: "драй"}
		expected := map[string]int64{"адын": 1, "цвай": 2, "драй": 3}
		res, err := SwapKeyValue(data)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty map passed", func(t *testing.T) {
		data := map[int64]string{}
		_, err := SwapKeyValue(data)
		require.Error(t, err)
	})

	t.Run("nil value passed", func(t *testing.T) {
		_, err := SwapKeyValue(nil)
		require.Error(t, err)
	})
}

func TestSliceFilter(t *testing.T) {
	t.Run("filter values that exist in base slice", func(t *testing.T) {
		data := []int64{1, 2, 3}
		filter := []int64{1, 2}
		expected := []int64{3}
		res, err := SliceFilter(data, filter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("filter values that do not exist in base slice", func(t *testing.T) {
		data := []int64{3, 4, 5, 6}
		filter := []int64{1, 2}
		expected := []int64{3, 4, 5, 6}
		res, err := SliceFilter(data, filter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("filter values in random order", func(t *testing.T) {
		data := []int64{1, 2, 3, 4, 5, 6}
		filter := []int64{2, 5}
		expected := []int64{1, 3, 4, 6}
		res, err := SliceFilter(data, filter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("nil first argument passed", func(t *testing.T) {
		filter := []int64{1, 2}
		_, err := SliceFilter(nil, filter)
		require.Error(t, err)
	})

	t.Run("nil second argument passed", func(t *testing.T) {
		data := []int64{1, 2}
		_, err := SliceFilter(data, nil)
		require.Error(t, err)
	})

	t.Run("filtered all values", func(t *testing.T) {
		data := []int64{1, 2, 3}
		filter := []int64{1, 2, 3}
		expected := []int64{}
		res, err := SliceFilter(data, filter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})
}

func TestSliceToMap(t *testing.T) {
	t.Run("cloned slice to map", func(t *testing.T) {
		data := []int64{1, 2, 3, 4}
		expected := map[int64]struct{}{int64(1): {}, int64(2): {}, int64(3): {}, int64(4): {}}
		res, err := SliceToMap(data)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("repeated values", func(t *testing.T) {
		data := []int64{1, 2, 2, 3}
		expected := map[int64]struct{}{int64(1): {}, int64(2): {}, int64(3): {}}
		res, err := SliceToMap(data)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("nil value passed", func(t *testing.T) {
		_, err := SliceToMap(nil)
		require.Error(t, err)
	})
}
