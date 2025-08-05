package main

import (
	"reflect"
)

// IsEqual compares two values of any type and returns true if they are equal, return false otherwise.
func IsEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	// Use reflection to determine the type
	typeOfA := reflect.TypeOf(a)
	typeOfB := reflect.TypeOf(b)

	// If types are different, values are not equal
	if typeOfA != typeOfB {
		return false
	}

	// Based on the type, use the appropriate comparison
	switch a.(type) {
	// Number types
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return reflect.DeepEqual(a, b)
	// String type
	case string:
		return a.(string) == b.(string)
	// Boolean type
	case bool:
		return a.(bool) == b.(bool)
	// For complex types, use reflect.DeepEqual
	case []interface{}, map[string]interface{}, struct{}:
		return reflect.DeepEqual(a, b)
	default:
		// For all other types, use reflect.DeepEqual
		return reflect.DeepEqual(a, b)
	}
}

// EqualNumbers compares two numbers regardless of their specific numeric type.
func EqualNumbers(a, b interface{}) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// Check if both values are numeric types
	if !isNumeric(va) || !isNumeric(vb) {
		return false
	}

	// Convert both to float64 for comparison
	var fa, fb float64

	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fa = float64(va.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fa = float64(va.Uint())
	case reflect.Float32, reflect.Float64:
		fa = va.Float()
	default:
		return false
	}

	switch vb.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fb = float64(vb.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fb = float64(vb.Uint())
	case reflect.Float32, reflect.Float64:
		fb = vb.Float()
	default:
		return false
	}

	return fa == fb
}

// EqualStrings compares two strings.
func EqualStrings(a, b string) bool {
	return a == b
}

// EqualBooleans compares two booleans.
func EqualBooleans(a, b bool) bool {
	return a == b
}

// EqualSlices compares two slices of any type
func EqualSlices(a, b interface{}) bool {
	sliceA := reflect.ValueOf(a)
	sliceB := reflect.ValueOf(b)

	// Check if both are slices
	if sliceA.Kind() != reflect.Slice || sliceB.Kind() != reflect.Slice {
		return false
	}

	// Check if lengths are equal
	if sliceA.Len() != sliceB.Len() {
		return false
	}

	// Compare each element
	for i := 0; i < sliceA.Len(); i++ {
		if !IsEqual(sliceA.Index(i).Interface(), sliceB.Index(i).Interface()) {
			return false
		}
	}

	return true
}

// EqualMaps compares two maps of any type.
func EqualMaps(a, b interface{}) bool {
	mapA := reflect.ValueOf(a)
	mapB := reflect.ValueOf(b)

	if mapA.Kind() != reflect.Map || mapB.Kind() != reflect.Map {
		return false
	}

	if mapA.Len() != mapB.Len() {
		return false
	}

	for _, key := range mapA.MapKeys() {
		valueA := mapA.MapIndex(key)
		valueB := mapB.MapIndex(key)

		// If key doesn't exist in map B
		if !valueB.IsValid() {
			return false
		}

		// Compare values for this key
		if !IsEqual(valueA.Interface(), valueB.Interface()) {
			return false
		}
	}

	return true
}

// EqualStructs compares two structs of any type.
func EqualStructs(a, b interface{}) bool {
	// For structs, we'll use reflect.DeepEqual which handles comparison of all fields recursively
	return reflect.DeepEqual(a, b)
}

// isNumeric checks if a reflect.Value is a numeric type
func isNumeric(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}
