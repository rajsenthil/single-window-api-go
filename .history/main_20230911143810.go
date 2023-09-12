package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/invopop/jsonschema"
)

type Product struct {
	ProductName string `json:"productName" jsonschema:"required"`
	UpcCode     string `json:"upcCode" jsonschema:"required,minLength=12,maxLength=12,description=Upc code,title=Upc Code,example=123456789012"`
	Attr01      string `json:"attr01"  jsonschema:"required"`
}

func main() {
	// s := jsonschema.Reflect(&Product{})
	// data, err := json.MarshalIndent(s, "", "  ")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(string(data))

	sch, err := jsonschema.Compile("testdata/person_schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := ioutil.ReadFile("testdata/person.json")
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
