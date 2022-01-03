package fxgo

import (
	"reflect"
	"strings"
	"testing"
)

func TestMapFilterReduce(t *testing.T) {
	var array = []int32{1, 2, 3, 4, 5}

	multiplyBy2 := func(x int32) int32 {
		return x * 2
	}

	sum := func(x int32, y int32) int32 {
		return x + y
	}

	dividByTwo := func(x int32) bool {
		return x%2 > 0
	}

	result :=
		Reduce(
			Map(
				Filter(array, dividByTwo), multiplyBy2), int32(0), sum)

	if reflect.ValueOf(result).Interface().(int32) != 18 {
		t.Error()
	}
}

func TestMapFilterReduce2(t *testing.T) {
	type People struct {
		Name string
		Age  int
	}

	peoples := []People{
		{Age: 13, Name: "hello"},
		{Age: 14, Name: "world"},
		{Age: 98, Name: "hololo"},
	}

	filterByAge := func(x People) bool {
		return x.Age <= 30
	}

	withOutOs := func(x People) People {
		return People{
			strings.ReplaceAll(x.Name, "o", ""),
			x.Age,
		}
	}
	result := Map(
		Filter(peoples, filterByAge), withOutOs)
	if !reflect.DeepEqual(result, []People{{Name: "hell", Age: 13}, {Name: "wrld", Age: 14}}) {
		t.Error()
	}
}
