package cry

import "testing"



type SomeStruct struct {
	Field1 string
	Field2 int
}



func Benchmark_Gob(b *testing.B) {
	temp := &SomeStruct{
		Field1: "Lorem ipsum",
		Field2: 12423,
	}
	for i := 0; i < b.N; i++ {
		data, _ := GobEncode(temp)
		op := new(SomeStruct)
		GobDecode(data, op)
	}
}


func Benchmark_Json(b *testing.B) {
	temp := &SomeStruct{
		Field1: "Lorem ipsum",
		Field2: 12423,
	}
	for i := 0; i < b.N; i++ {
		data, _ := JsonEncode(temp)
		op := new(SomeStruct)
		JsonDecode(data, op)
	}
}
