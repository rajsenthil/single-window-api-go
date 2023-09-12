package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/santhosh-tekuri/jsonschema"
)

type Product struct {
	ProductName string `json:"productName"`
	UpcCode     string `json:"upcCode"`
	Attr01      string `json:"attr01"`
}

func main() {
	sch, err := jsonschema.Compile("schema/product-attr-validation-schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := ioutil.ReadFile("schema/product01.json")
	if err != nil {
		log.Fatal(err)
	}

	// var prod = new(Product)
	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		log.Fatal(err)
	}

	log.Printf("Prod1: %v", v)
	fmt.Printf("Prod2: %v", v)
	if err = sch.ValidateInterface(v); err != nil {
		log.Fatalf("Validation Error %v", err)
	}
	fmt.Printf("Prod3: %v", v)
}
