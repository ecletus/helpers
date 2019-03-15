package helpers

import (
	"reflect"
)

var NILL_INTERFACE interface{}
var NILL_INTERFACE_VALUE = reflect.ValueOf(NILL_INTERFACE)

// check if value is nil
func IsNil(value reflect.Value) bool {
	if value.Kind() != reflect.Ptr {
		return false
	}
	if value.Pointer() == 0 {
		return true
	}
	v := value.Elem().Interface()
	if v == NILL_INTERFACE {
		return true
	}
	return false
}

func IsNilInterface(value interface{}) bool {
	if value == nil {
		return true
	}
	return IsNil(reflect.ValueOf(value))
}

func StringsToInterfaces(values []string) (r []interface{}) {
	for _, s := range values {
		r = append(r, s)
	}
	return
}