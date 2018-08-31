package json

import "github.com/minamijoyo/tfschema/tfschema"

// Attribute is wrapper for tfschema.Attribute.
type Attribute struct {
	// Name is a name of attribute.
	// Note that this key does not exist in the original data structure.
	// In order to reduce the possibility of future conflicts,
	// naming borrowed from the schema definition of the grpc provider's proto file.
	// https://github.com/hashicorp/terraform/pull/18550
	Name string `json:"name"`

	// Type is a type of the attribute's value.
	// Note that Type is not cty.Type
	// We cannot import github.com/hashicorp/terraform/vendor/github.com/zclconf/go-cty/cty
	// On the other hand, tfschema does not need a dynamic type.
	// So, we use a simple representation of type defined in tfschema package.
	Type tfschema.Type `json:"type"`
	// Required is a flag whether this attribute is required.
	Required bool `json:"required"`
	// Optional is a flag whether this attribute is optional.
	// This field conflicts with Required.
	Optional bool `json:"optional"`
	// Computed is a flag whether this attribute is computed.
	// If true, the value may come from provider rather than configuration.
	// If combined with Optional, then the config may optionally provide an
	// overridden value.
	Computed bool `json:"computed"`
	// Sensitive is a flag whether this attribute may contain sensitive information.
	Sensitive bool `json:"sensitive"`
}

// NewAttribute creates a new Attribute instance.
func NewAttribute(a *tfschema.Attribute, name string) *Attribute {
	return &Attribute{
		Name:      name,
		Type:      a.Type,
		Required:  a.Required,
		Optional:  a.Optional,
		Computed:  a.Computed,
		Sensitive: a.Sensitive,
	}
}
