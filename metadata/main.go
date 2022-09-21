package main

import (
	"fmt"
	"github.com/deislabs/metadata/metadata"
	"time"
)

type Metadata struct {
	Mystring          string            `json:"mystring"`
	Myduration        metadata.Duration `json:"myduration"`
	Myinteger         int               `json:"myinteger,string"`
	Myfloat64         float64           `json:"myfloat64,string"`
	Mybool            *bool             `json:"mybool,omitempty"`
	MyRegularDuration time.Duration     `json:"myregularduration"`
}

func main() {

	var m Metadata

	data := make(map[string]string)
	data["mystring"] = "test"
	data["myduration"] = "3s"
	data["myinteger"] = "1"
	data["myfloat64"] = "1.1"
	data["mybool"] = "true"
	data["myregularduration"] = "6m"

	err := metadata.DecodeMetadata(data, &m)

	if err != nil {
		fmt.Println("decode metadata errored out ", err)
	}

	value, err := metadata.IsRawPayload(data)
	if err != nil {
		fmt.Println("IsRawPayload errored out ", err)
	}

	fmt.Println("Is Raw Payload ", value)
}
