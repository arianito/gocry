package cry

import "testing"



type SomeStruct struct {
	Field1 string
	Field2 int
}



func Benchmark_Serialize_Deserialize(b *testing.B) {
	temp := &SomeStruct{
		Field1: "Lorem ipsum",
		Field2: 12423,
	}
	for i := 0; i < b.N; i++ {
		data, _ := Serialize(temp)

		op := new(SomeStruct)
		Deserialize(data, op)
	}
}
