package main

import (
	"bytes"
	"encoding/json"
	"os"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	f, err := os.OpenFile("json.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// f.Write([]byte("writing some string into the file\n"))
	// f.Write([]byte("writing some string into the file\n"))

	buff := new(bytes.Buffer)
	enc := json.NewEncoder(buff)

	u := user{
		Name: "bob",
		Age:  11,
	}

	if err := enc.Encode(u); err != nil {
		panic(err)
	}

	// {"name":"bob","age":10}

	f.Write(buff.Bytes())

}
