package main

import (
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type Product struct {
	ProductName string `json:"productName" jsonschema:"required"`
	UpcCode     string `json:"upcCode" jsonschema:"required,minLength=12,maxLength=12,description=Upc code,title=Upc Code,example=123456789012"`
	Attr01      string `json:"attr01"  jsonschema:"required"`
}

func main() {
	s := jsonschema.Reflect(&Product{})
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
}
