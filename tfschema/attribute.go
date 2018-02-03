package tfschema

import (
	"github.com/hashicorp/terraform/config/configschema"
)

// Attribute is wrapper for configschema.Attribute
type Attribute struct {
	// Type is a type of the attribute's value.
	// Note that Type is not cty.Type
	// We cannot import github.com/hashicorp/terraform/vendor/github.com/zclconf/go-cty/cty
	// On the other hand, tfschema does not need a dynamic type.
	// So, we use a simple representation of type.
	Type      Type `json:"type"`
	Required  bool `json:"required"`
	Optional  bool `json:"optional"`
	Computed  bool `json:"computed"`
	Sensitive bool `json:"sensitive"`
}

func NewAttribute(a *configschema.Attribute) *Attribute {
	return &Attribute{
		Type:      *NewType(a.Type),
		Required:  a.Required,
		Optional:  a.Optional,
		Computed:  a.Computed,
		Sensitive: a.Sensitive,
	}
}
