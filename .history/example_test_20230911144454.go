package jsonschema_test

import (
	"encoding/json"
	"log"
	"os"

	"github.com/santhosh-tekuri/jsonschema"
)

func Example() {
	sch, err := jsonschema.Compile("testdata/person_schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := os.ReadFile("testdata/person.json")
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
	// Output:
}
