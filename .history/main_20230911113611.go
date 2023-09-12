package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/santhosh-tekuri/jsonschema"
)

type product struct {
	ProductName string `json:"productName"`
	UpcCode     int    `json:"upcCode"`
	Attr01      string `json:"attr01"`
}

// Read implements io.Reader.
func (product) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

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

	if err = sch.Validate(prod); err != nil {
		log.Fatalf("%#v", err)
	}
}