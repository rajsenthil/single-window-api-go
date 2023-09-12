package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/santhosh-tekuri/jsonschema"
)

type Product struct {
	ProductName string `json:"productName"`
	UpcCode     int    `json:"upcCode"`
	Attr01      string `json:"attr01"`
}

// Read implements io.Reader.
func (Product) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

func main() {
	sch, err := jsonschema.Compile("testdata/person_schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := ioutil.ReadFile("testdata/person.json")
	if err != nil {
		log.Fatal(err)
	}

	var v interface{}
	var prod = Product{
		ProductName: "",
		UpcCode:     123456789012,
		Attr01:      "Attr01",
	}
	if err := json.Unmarshal(data, &prod); err != nil {
		log.Fatal(err)
	}

	if err = sch.Validate(prod); err != nil {
		log.Fatalf("%#v", err)
	}

}