package models

import "sort"

import "../../terraform-provider-aws/aws"

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

func ResourceType(i aws.TerraformResourceType) string {
	values := []string{
		"DataSource",
		"Resource",
	}

	return values[i]
}

func ResourceBuilder(terraformResource aws.TerraformResource) Resource {
	schemas := []Schema{}
	for name, schema := range terraformResource.Value.Schema {
		schemas = append(
			schemas,
			SchemaBuilder(name, schema),
		)
	}

	sort.Slice(schemas, func(i, j int) bool {
		return schemas[i].Name < schemas[j].Name
	})

	return Resource{
		Type: ResourceType(terraformResource.Type),
		Name: terraformResource.Name,
		Value: ResourceValue{
			Schema: schemas,
			SchemaVersion: terraformResource.Value.SchemaVersion,
			DeprecationMessage: terraformResource.Value.DeprecationMessage,
		},
	}
}
