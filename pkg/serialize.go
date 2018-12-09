package cry

import (
	"bytes"
	"encoding/gob"
)

// Serialize any struct to blob of data
func Serialize(a interface{}) ([]byte, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(a)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Deserialize blob of data to struct
func Deserialize(a []byte, o interface{}) error {
	b := bytes.Buffer{}
	_, err := b.Write(a)
	if err != nil {
		return err
	}
	d := gob.NewDecoder(&b)
	return d.Decode(o)
}
