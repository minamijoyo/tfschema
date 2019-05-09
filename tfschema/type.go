package tfschema

import (
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
// This method customize cty.GoString() to make it easy to read.
func (t *Type) Name() (string, error) {
	goString := t.GoString()
	// drop `cty.` prefix for simplicity. (e.g. cty.String => String)
	name := strings.Replace(goString, "cty.", "", -1)

	return name, nil
}
