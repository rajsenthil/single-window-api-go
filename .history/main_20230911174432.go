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

	var prod = Product{}

	if err := json.Unmarshal(data, &prod); err != nil {
		log.Fatal(err)
	}

	log.Printf("Prod: %v", prod)
	fmt.Printf("Prod: %v", prod)
	if err = sch.ValidateInterface(prod); err != nil {
		log.Fatalf("Error %v", err)
	}
}