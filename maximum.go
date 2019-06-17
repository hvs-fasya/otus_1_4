package otus_1_4

import (
	"reflect"
)

//GetMax return maximim elelemnt af arr slice using comparator-function less
func GetMax(input interface{}, less func(i, j int) bool) interface{} {
	rv := reflect.ValueOf(input)
	if reflect.TypeOf(input).Kind() != reflect.Slice || rv.Len() == 0 {
		return nil
	}
	var maxIndex int
	for i := 0; i < rv.Len(); i++ {
		if less(maxIndex, i) {
			maxIndex = i
		}
	}
	return rv.Index(maxIndex).Interface()
}
