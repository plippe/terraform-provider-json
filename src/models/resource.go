package models

import "sort"

import "../../terraform-provider-aws/aws"
import "github.com/hashicorp/terraform/helper/schema"

type Resource struct {
	Type  string
	Name  string
	Value ResourceValue
}

type ResourceValue struct {
	Schema             []Schema
	SchemaVersion      int
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
	return Resource{
		Type:  ResourceType(terraformResource.Type),
		Name:  terraformResource.Name,
		Value: ResourceValueBuilder(terraformResource.Value),
	}
}

func ResourceValueBuilder(terraformResource *schema.Resource) ResourceValue {
	schemas := []Schema{}
	for name, schema := range terraformResource.Schema {
		schemas = append(
			schemas,
			SchemaBuilder(name, schema),
		)
	}

	sort.Slice(schemas, func(i, j int) bool {
		return schemas[i].Name < schemas[j].Name
	})

	return ResourceValue{
		Schema:             schemas,
		SchemaVersion:      terraformResource.SchemaVersion,
		DeprecationMessage: terraformResource.DeprecationMessage,
	}
}
