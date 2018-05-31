package main

import (
	"testing"

)

type  testpairs struct {
	value,output int
}
var testingarray =[]testpairs  {{1,2},{50,64},{3,4}}

func TestTask(t *testing.T) {
	for _,value:=range testingarray{
		temp:= Task(value.value)
		if value.output!= temp {
			t.Error("For ",value.value, "expected", value.output, "got", temp)
		}
	}
}
