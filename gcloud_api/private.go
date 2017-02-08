package gcloud_api

import (
	"bytes"
	"encoding/gob"
)

func makeBytes(contents interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(contents)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

