package main

import (
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type Pet struct {
	// Name of the animal.
	Name string `json:"name" jsonschema:"title=Name"`
}

// Pets is a collection of Pet objects.
type Pets []*Pet

// NamedPets is a map of animal names to pets.
type NamedPets map[string]*Pet

type (
	// Plant represents the plants the user might have and serves as a test
	// of structs inside a `type` set.
	Plant struct {
		Variant string `json:"variant" jsonschema:"title=Variant"` // This comment will be ignored
	}
)

func nested_schema_generator() {
	s := jsonschema.Reflect(&NamedPets{})
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
}
