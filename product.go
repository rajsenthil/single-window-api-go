package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/invopop/jsonschema"
	jsonschema2 "github.com/qri-io/jsonschema"
)

type Product struct {
	Id           int32                  `json:"id" jsonschema:"required"`
	ProductType  ProductType            `json:"productType" jsonschema:"required,minLength=1,maxLength=20,enum=Simple,enum=Configurable,enum=Grouped,enum=Virtual,enum=Bundle,title=product type,description=product type,example=Bundle,example=Configurable,default=Simple"`
	Product_date time.Time              `json:"product_date" jsonschema:"required" jsonschema_description:"Product date must be valid" jsonschema_message:"Product date is required and date time must be valid"`
	ProductPrice ProductPrice           `json:"productPrice" jsonschema:"required"`
	Tags         map[string]interface{} `json:"tags,omitempty" jsonschema_extras:"a=b,foo=bar,foo=bar1"`
}

type ProductPrice struct {
	RegularPrice float32 `json:"regularPrice"`
	MarkedPrice  float32 `json:"markedPrice"`
	Discount     int     `json:"discount,omitempty"`
}

type ProductType string

func product_schema_generator() {
	s := jsonschema.Reflect(&Product{})
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
}

func validate_product() {
	ctx := context.Background()
	sch, err := os.ReadFile("schema/product-schema.json")
	if err != nil {
		log.Fatal(err)
	}

	rs := &jsonschema2.Schema{}
	if err := json.Unmarshal(sch, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	product01, err := os.ReadFile("schema/product.json")
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
