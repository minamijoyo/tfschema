package tfschema

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zclconf/go-cty/cty"
)

// Type is a type of the attribute's value.
type Type struct {
	// We embed cty.Type to customize string representation.
	cty.Type
}

// NewType creates a new Type instance.
func NewType(t cty.Type) *Type {
	return &Type{
		Type: t,
	}
}

// MarshalJSON returns a encoded string in JSON.
func (t *Type) MarshalJSON() ([]byte, error) {
	name, err := t.Name()
	if err != nil {
		return nil, err
	}

	return []byte(`"` + name + `"`), nil
}

// Name returns a name of type.
// Terraform v0.12 introduced a new `SchemaConfigModeAttr` feature.
// Most attributes have simple types, but if `SchemaConfigModeAttr` is set for
// an attribute, it is syntactically NestedBlock but semantically interpreted
// as an Attribute. In this case, Attribute has a complex data type. It is
// reasonable to use the same notation as the type annotation in HCL2 to
// represent the correct data type. However, it seems that HCL2 has a type
// annotation parser but no writer, so we implement it by ourselves.
//
// See also:
// - https://github.com/minamijoyo/tfschema/issues/9
// - https://github.com/terraform-providers/terraform-provider-aws/pull/8187
// - https://github.com/hashicorp/terraform/pull/20626
// - https://www.terraform.io/docs/configuration/types.html
func (t *Type) Name() (string, error) {
	switch {
	case t.IsPrimitiveType():
		switch t.Type {
		case cty.String:
			return "string", nil
		case cty.Number:
			return "number", nil
		case cty.Bool:
			return "bool", nil
		}

	case t.IsListType():
		elementType := NewType(t.ElementType())
		elementName, err := elementType.Name()
		if err != nil {
			return "", err
		}
		return "list(" + elementName + ")", nil

	case t.IsSetType():
		elementType := NewType(t.ElementType())
		elementName, err := elementType.Name()
		if err != nil {
			return "", err
		}
		return "set(" + elementName + ")", nil

	case t.IsMapType():
		elementType := NewType(t.ElementType())
		elementName, err := elementType.Name()
		if err != nil {
			return "", err
		}
		return "map(" + elementName + ")", nil

	case t.IsTupleType():
		elementTypes := t.TupleElementTypes()
		elementNames := make([]string, 0, len(elementTypes))
		for i := range elementTypes {
			elementType := NewType(t.TupleElementType(i))
			elementName, err := elementType.Name()
			if err != nil {
				return "", err
			}

			elementNames = append(elementNames, elementName)
		}
		return "tuple([ " + strings.Join(elementNames, ", ") + " ])", nil
	case t.IsObjectType():
		attributeTypes := t.AttributeTypes()
		attributeNames := make([]string, 0, len(attributeTypes))
		for k := range attributeTypes {
			attributeNames = append(attributeNames, k)
		}
		sort.Strings(attributeNames)

		attributes := make([]string, 0, len(attributeTypes))
		for _, k := range attributeNames {
			elementType := NewType(t.AttributeType(k))
			elementName, err := elementType.Name()
			if err != nil {
				return "", err
			}

			attributes = append(attributes, k+"="+elementName)
		}
		return "object({ " + strings.Join(attributes, ", ") + " })", nil

	case t.IsCapsuleType():
		// We notice that there is a capsule type as a cty specification,
		// but we don't know how to handle it properly,
		// so we should make an error for now..
		return "", fmt.Errorf("Failed to get name for unsupported capsule type: %#v", t)

	default:
		// should never happen
		return "", fmt.Errorf("Failed to get name for unknown type: %#v", t)
	}

	// should never happen
	return "", fmt.Errorf("Failed to get name for type: %#v", t)
}
