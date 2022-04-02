package arrayutils

// Array - Encapsulate array type in a generic struct
//
// This struct is useful for chaining methods and doing "functional" programming
//
// Ex:
//  myarr := Array[int]{[]int{1, 2, 3, 4, 5}}
//  result := myarr.Map(func(val int) int {
//    return val * 3
//  }).Filter(func(val int) bool {
//    return val % 2 == 0
//  })
// => result [6 12]
type Array[T any] struct {
	Arr []T
}

// Find - Find a value in the current array
func (a *Array[T]) Find(f func(T) bool) T {
	var notfound T
	for _, val := range a.Arr {
		if f(val) {
			return val
		}
	}
	return notfound
}

// FindIndex - Find value index in the current array
func (a *Array[T]) FindIndex(f func(T) bool) int {
	for i, val := range a.Arr {
		if f(val) {
			return i
		}
	}
	return -1
}

// Filter - Filter array values (chainable)
func (a *Array[T]) Filter(f func(T) bool) *Array[T] {
	var filtered []T
	for _, val := range a.Arr {
		if f(val) {
			filtered = append(filtered, val)
		}
	}
	return &Array[T]{filtered}
}

// Map - Create a new array populated with the result of calling the input func on each array element (chainable)
func (a *Array[T]) Map(f func(T) T) *Array[T] {
	var filtered []T
	for _, val := range a.Arr {
		filtered = append(filtered, f(val))
	}
	return &Array[T]{filtered}
}
