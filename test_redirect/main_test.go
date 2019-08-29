package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestDDD(t *testing.T) {
	buf,_:=json.Marshal(map[string]interface{}{
		"hello":"1",
	})
	http.Post("http://localhost:8808/1", "application/json", bytes.NewReader(buf))
}
