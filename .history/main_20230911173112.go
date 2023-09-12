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

// Read implements io.Reader.
// func (Product) Read(p []byte) (n int, err error) {
// 	// panic("unimplemented")
// }

func main() {
	sch, err := jsonschema.Compile("schema/product-attr-validation-schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := ioutil.ReadFile("schema/product01.json")
	if err != nil {
		log.Fatal(err)
	}

	// var v interface{}
	var prod = Product{
		ProductName: "Test",
		UpcCode:     "123456789012",
		Attr01:      "Attr01",
	}
	if err := json.Unmarshal(data, &prod); err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", prod)
	fmt.Printf("%v", prod)
	err = sch.ValidateInterface(prod)
	log.Printf("%v", err)
	if err != nil {
		log.Fatalf("Error %#v", err)
	}

}