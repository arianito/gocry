package cry

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// Serialize any struct to blob of data
func GobEncode(a interface{}) ([]byte, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(a)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Deserialize blob of data to struct
func GobDecode(a []byte, o interface{}) error {
	b := bytes.Buffer{}
	_, err := b.Write(a)
	if err != nil {
		return err
	}
	d := gob.NewDecoder(&b)
	return d.Decode(o)
}



func JsonEncode(a interface{}) ([]byte, error) {
	return json.Marshal(a)
}

func JsonDecode(a []byte, o interface{}) error {
	return json.Unmarshal(a, o)
}
