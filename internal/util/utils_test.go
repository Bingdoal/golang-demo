package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	var test = []int{1, 5, 7, 8, 9, 3}
	var result = Filter(test, func(e int) bool {
		return e <= 3
	})
	fmt.Printf("result: %v\n", result)
	assert.Equal(t, result, []int{1, 3})
}

func TestMap(t *testing.T) {
	var test = []int{1, 5, 7, 8, 9, 3}
	var result = Map(test, func(e int) int {
		return e * 2
	})
	fmt.Printf("result: %v\n", result)
	assert.Equal(t, result, []int{2, 10, 14, 16, 18, 6})
}

func TestNewPromise(t *testing.T) {
	var promise = NewPromise(func() (int, error) {
		var i = 1 + 1
		return i, nil
	})

	promise.Then(func(i int) {
		fmt.Printf("then i: %v\n", i)
	})
	promise.Wait()
	fmt.Printf("then end\n")

	promise = NewPromise(func() (int, error) {
		return 0, fmt.Errorf("error")
	})

	promise.Then(func(i int) {
		fmt.Printf("then i: %v\n", i)
	}).Catch(func(e error) {
		fmt.Printf("catch e: %v\n", e)
	})
	promise.Wait()
	fmt.Printf("catch end\n")
}
