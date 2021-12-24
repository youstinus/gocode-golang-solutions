// Objective: Start the Jet Ski

package main

import "reflect"

func main() {
	println("Logging in...")
	authorized := startup(login())
	if reflect.ValueOf(authorized).Bool() {
		println("Starting the engine")
		return
	}
	println("Startup failed")
}

func validSequence(i int, el interface{}) bool {
	return reflect.TypeOf(el).String() == "*main.Sequence" &&
		!reflect.ValueOf(el).IsNil() &&
		reflect.ValueOf(el).Elem().NumField() == 2 &&
		reflect.TypeOf(reflect.ValueOf(el).Elem().Field(0).Interface()).String() == "int" &&
		int(reflect.ValueOf(el).Elem().Field(0).Int()) == i*i-i &&
		!reflect.ValueOf(reflect.ValueOf(el).Elem().Field(1).Interface()).IsNil()
}

func startup(seq interface{}) bool {
	for i := 0; i < 5; i++ {
		if !validSequence(i, seq) {
			return false
		}
		seq = reflect.ValueOf(seq).Elem().Field(1).Interface()
	}

	return true
}


func login() *Sequence {
	return &Sequence{0, &Sequence{0, &Sequence{2, &Sequence{6, &Sequence{12, &Sequence{20, &Sequence{30, &Sequence{}}}}}}}}
}

type Sequence struct {
	Age  int
	Name *Sequence
}

func (s *Sequence) Interface() *Sequence {
	return s
}