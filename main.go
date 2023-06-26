package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/linkedin/goavro/v2"
	"go.einride.tech/protobuf-avro/protoavro"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	// Parse command line arguments
	jsonFile := flag.String("f", "", "Input JSON Protobuf schema file")
	flag.Parse()

	// Read JSON file
	jsonBytes, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Unmarshal JSON into a protobuf descriptor
	var desc descriptorpb.DescriptorProto
	if err := json.Unmarshal(jsonBytes, &desc); err != nil {
		log.Fatalf("Failed to unmarshal JSON schema: %v", err)
	}

	// Infer Avro schema
	schema, err := protoavro.InferSchema(desc.ProtoReflect().Descriptor())
	if err != nil {
		log.Fatalf("Failed to infer Avro schema: %v", err)
	}

	// Convert schema to JSON
	codec, err := goavro.NewCodec(schema.String())
	if err != nil {
		log.Fatalf("Failed to create Avro codec: %v", err)
	}

	// Print Avro schema to console
	fmt.Println(codec.Schema())
}
