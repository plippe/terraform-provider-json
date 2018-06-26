package main

import "fmt"
import "os"
import "sort"
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

func cast(terraformResource aws.TerraformResource) Resource {
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

	sort.Slice(schemas, func(i, j int) bool {
		return schemas[i].Name < schemas[j].Name
	})

	return Resource{
		Type: ResourceTypes()[terraformResource.Type],
		Name: terraformResource.Name,
		Value: ResourceValue{
			Schema: schemas,
			SchemaVersion: terraformResource.Value.SchemaVersion,
			DeprecationMessage: terraformResource.Value.DeprecationMessage,
		},
	}
}

func write(f *os.File, resource Resource) error {
	bytes, err := json.Marshal(resource)
	if err == nil { _, err = f.Write(bytes) }
	if err == nil { _, err = f.Write([]byte("\n")) }

	return err
}

func main() {
	terraformResources := append(aws.DataSources(), aws.Resources()...)

	resources := []Resource{}
	for _, terraformResource := range terraformResources {
		newResource := cast(terraformResource)
		resources = append(resources, newResource)
	}

	filePath := "docs/terraform-provider-aws.json"

	err := os.Remove(filePath)
	if err != nil { panic(err) }

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil { panic(err) }
    defer f.Close()

	for _, resource := range resources {
		err = write(f, resource)
		if err != nil { panic(err) }
	}

	fmt.Println("done")
}
