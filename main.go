package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/qri-io/jsonschema"
)

type Product struct {
	ProductName string `json:"productName" jsonschema:"required"`
	UpcCode     string `json:"upcCode" jsonschema:"required,minLength=12,maxLength=12,description=Upc code,title=Upc Code,example=123456789012"`
	Attr01      string `json:"attr01"  jsonschema:"required"`
}

func main() {
	ctx := context.Background()
	sch, err := os.ReadFile("schema/product-attr-validation-schema.json")
	if err != nil {
		log.Fatal(err)
	}

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(sch, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	product01, err := os.ReadFile("schema/product01.json")
	if err != nil {
		log.Fatal(err)
	}

	errs, err := rs.ValidateBytes(ctx, product01)
	if err != nil {
		panic(err)
	}

	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
	}

}
