package data

import (
	"bytes"
	"encoding/gob"
)

// InterfaceToBytes convert a interface{} to a []byte
func InterfaceToBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer

	//message.WriteDebug("(InterfaceToBytes)")
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
