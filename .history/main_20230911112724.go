package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/santhosh-tekuri/jsonschema"
)

type product struct {
	ProductName string `json:"productName"`
	UpcCode     int16  `json:"upcCode`
	attr01      string
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

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		log.Fatal(err)
	}

	if err = sch.Validate(v); err != nil {
		log.Fatalf("%#v", err)
	}
}
