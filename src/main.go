package main

import "fmt"
import "encoding/json"

import "../terraform-provider-aws/aws"

type Resource struct {
	Type string
	Name string
	Value ResourceValue
}

type ResourceValue struct {
	Schema []Schema
	SchemaVersion int
	DeprecationMessage string
}

type Schema struct {
	Name string
	Value SchemaValue
}

type SchemaValue struct {
	Type string
	Optional bool
	Required bool
	Default interface{}
	Description string
	InputDefault string
	Computed bool
	ForceNew bool
	// Elem interface{}
	MaxItems int
	MinItems int
	PromoteSingle bool
	ComputedWhen []string
	ConflictsWith []string
	Deprecated string
	Removed string
	Sensitive bool
}

func ResourceTypes() []string {
	return []string{
		"DataSource",
		"Resource",
	}
}

func Types() []string {
	return []string{
		"Invalid",
		"Bool",
		"Int",
		"Float",
		"String",
		"List",
		"Map",
		"Set",
		"Object",
	}
}

func main() {
	terraformResources := append(aws.DataSources(), aws.Resources()...)

	resources := []Resource{}
	for _, terraformResource := range terraformResources {

		schemas := []Schema{}
		for name, schema := range terraformResource.Value.Schema {
			newSchema := Schema{
				Name: name,
				Value: SchemaValue{
					Type: Types()[schema.Type],
					Optional: schema.Optional,
					Required: schema.Required,
					Default: schema.Default,
					Description: schema.Description,
					InputDefault: schema.InputDefault,
					Computed: schema.Computed,
					ForceNew: schema.ForceNew,
					// Elem: schema.Elem,
					MaxItems: schema.MaxItems,
					MinItems: schema.MinItems,
					PromoteSingle: schema.PromoteSingle,
					ComputedWhen: schema.ComputedWhen,
					ConflictsWith: schema.ConflictsWith,
					Deprecated: schema.Deprecated,
					Removed: schema.Removed,
					Sensitive: schema.Sensitive,
				},
			}

			schemas = append(schemas, newSchema)
		}

		newResource := Resource{
			Type: ResourceTypes()[terraformResource.Type],
			Name: terraformResource.Name,
			Value: ResourceValue{
				Schema: schemas,
				SchemaVersion: terraformResource.Value.SchemaVersion,
				DeprecationMessage: terraformResource.Value.DeprecationMessage,
			},
		}

		resources = append(resources, newResource)
	}

	for _, resource := range resources {
		bytes, err := json.Marshal(resource)
		if err != nil { panic(err) }

		fmt.Println(string(bytes))
	}
}
