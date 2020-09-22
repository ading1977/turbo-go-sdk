package main

import (
	"encoding/json"
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/dataingestionframework/data"
)

func main() {
	actualSchema := data.GenerateJSONSchema()
	actualJSON, _ := json.MarshalIndent(actualSchema, "", "  ")
	fmt.Println(string(actualJSON))
}

