package cry

import "testing"



type SomeStruct struct {
	Field1 string
	Field2 int
	Field3 map[string]string
}



func Benchmark_Serialize_Deserialize(b *testing.B) {
	temp := &SomeStruct{
		Field1: "hello world!",
		Field2: 12423,
		Field3: map[string]string{
			"are":     "you",
			"kidding": "me ?",
			"why ?":   "i'm exhausted writing this",
		},
	}
	for i := 0; i < b.N; i++ {
		data, _ := Serialize(temp)

		op := new(SomeStruct)
		Deserialize(data, op)
	}
}
