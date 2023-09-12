package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/santhosh-tekuri/jsonschema"
)

type product struct {
	ProductName string `json:"productName"`
	UpcCode     string `json:"upcCode"`
	Attr01      string `json:"attr01"`
}

// Read implements io.Reader.
func (product) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// func unmarshal(r io.Reader) (interface{}, error) {
// 	decoder := json.NewDecoder(r)
// 	decoder.UseNumber()
// 	var doc interface{}
// 	if err := decoder.Decode(&doc); err != nil {
// 		return nil, err
// 	}
// 	if t, _ := decoder.Token(); t != nil {
// 		return nil, fmt.Errorf("invalid character %v after top-level value", t)
// 	}
// 	return doc, nil
// }

func main() {
	print("Testing...")
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft7

	sch, err := jsonschema.Compile("schema/product-attr-validation-schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := os.ReadFile("schema/product01.json")
	if err != nil {
		log.Fatal(err)
	}

	// var v interface{}
	var prod = product{}
	if err := json.Unmarshal(data, &prod); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", prod)
	if err = sch.Validate(prod); err != nil {
		log.Fatalf("%#v", err)
	}
}
