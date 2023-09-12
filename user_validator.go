package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/qri-io/jsonschema"
)

func validate_user() {
	ctx := context.Background()
	sch, err := os.ReadFile("schema/user-schema.json")
	if err != nil {
		log.Fatal(err)
	}

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(sch, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	user01, err := os.ReadFile("schema/user.json")
	if err != nil {
		log.Fatal(err)
	}

	errs, err := rs.ValidateBytes(ctx, user01)
	if err != nil {
		panic(err)
	}

	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
	}

}
