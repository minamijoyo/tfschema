package tfschema

import (
	"fmt"
	"reflect"
	"strings"
)

// Type is a type of the attribute's value.
type Type struct {
	// T is an instance of github.com/hashicorp/terraform/vendor/github.com/zclconf/go-cty.Type
	// but we cannot import it, so we embed it here with the interface{}.
	ctyType interface{}
}

// NewType creates a new Type instance.
func NewType(t interface{}) *Type {
	return &Type{
		ctyType: t,
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
// This method depends on the private method of cty.typeImpl.GoString().
// It's fragile but in the meantime easy to implement.
// Ideally it should be implemented by looking at the type of cty.typeImpl.
func (t *Type) Name() (string, error) {
	v := reflect.ValueOf(t.ctyType).MethodByName("GoString")
	if !v.IsValid() {
		return "", fmt.Errorf("Faild to find GoString(): %#v", t)
	}

	nv := v.Call([]reflect.Value{})
	if len(nv) == 0 {
		return "", fmt.Errorf("Faild to call GoString(): %#v", v)
	}

	goString := nv[0].String()
	// drop `cty.` prefix for simplicity. (e.g. cty.String => String)
	name := strings.Replace(goString, "cty.", "", -1)

	return name, nil
}
