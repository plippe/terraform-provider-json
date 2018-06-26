package models

import "github.com/hashicorp/terraform/helper/schema"

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

func SchemaType(i schema.ValueType) string {
	values := []string{
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

	return values[i]
}

func SchemaBuilder(name string, terraformSchema *schema.Schema) Schema {
	return Schema{
		Name: name,
		Value: SchemaValue{
			Type: SchemaType(terraformSchema.Type),
			Optional: terraformSchema.Optional,
			Required: terraformSchema.Required,
			Default: terraformSchema.Default,
			Description: terraformSchema.Description,
			InputDefault: terraformSchema.InputDefault,
			Computed: terraformSchema.Computed,
			ForceNew: terraformSchema.ForceNew,
			// Elem: terraformSchema.Elem,
			MaxItems: terraformSchema.MaxItems,
			MinItems: terraformSchema.MinItems,
			PromoteSingle: terraformSchema.PromoteSingle,
			ComputedWhen: terraformSchema.ComputedWhen,
			ConflictsWith: terraformSchema.ConflictsWith,
			Deprecated: terraformSchema.Deprecated,
			Removed: terraformSchema.Removed,
			Sensitive: terraformSchema.Sensitive,
		},
	}
}
