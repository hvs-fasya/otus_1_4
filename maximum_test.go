package otus_1_4

import (
	"reflect"
	"testing"
)

type Aged struct {
	Age int
}

func TestGetMax_Empty(t *testing.T) {
	var empty = []string{}
	if got := GetMax(empty, func(i, j int) bool { return len([]byte(empty[i])) < len([]byte(empty[j])) }); !reflect.DeepEqual(got, nil) {
		t.Errorf("GetMax() = %+v, want %+v", got, nil)
	}
}

func TestGetMax_StringsByLen(t *testing.T) {
	var test = []string{"a", "bbb", "ccccc", "dd"}
	var want = test[2]
	if got := GetMax(test, func(i, j int) bool { return len([]byte(test[i])) < len([]byte(test[j])) }); !reflect.DeepEqual(got, want) {
		t.Errorf("GetMax() = %+v, want %+v", got, want)
	}
}

func TestGetMax_Struct(t *testing.T) {
	var arrAged = []Aged{
		{Age: 11},
		{Age: 15},
		{Age: 0},
	}
	var want = arrAged[1]
	if got := GetMax(arrAged, func(i, j int) bool { return arrAged[i].Age < arrAged[j].Age }); !reflect.DeepEqual(got, arrAged[1]) {
		t.Errorf("GetMax() = %+v, want %+v", got, want)
	}
}
