package utils

import "reflect"

// Map Example:
//   numbers := []int{1, 2, 3}
//   fmt.Println("Square", numbers, f.Map(func(x int) int { return x * x }, numbers))
//
// origin by: https://github.com/izqui/functional
func Map(f interface{}, vs interface{}) interface{} {
	vf := reflect.ValueOf(f)
	vx := reflect.ValueOf(vs)
	l := vx.Len()
	tys := reflect.SliceOf(vf.Type().Out(0))
	vys := reflect.MakeSlice(tys, l, l)
	for i := 0; i < l; i++ {
		y := vf.Call([]reflect.Value{vx.Index(i)})[0]
		vys.Index(i).Set(y)
	}
	return vys.Interface()
}
