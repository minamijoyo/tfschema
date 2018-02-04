package tfschema

import (
	"fmt"
	"reflect"
	"strings"
)

type Type struct {
	// T is an instance of github.com/hashicorp/terraform/vendor/github.com/zclconf/go-cty.Type
	ctyType interface{}
}

func NewType(t interface{}) *Type {
	return &Type{
		ctyType: t,
	}
}

func (t *Type) MarshalJSON() ([]byte, error) {
	name, err := t.Name()
	if err != nil {
		return nil, err
	}

	return []byte(`"` + name + `"`), nil
}

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
	name := strings.Replace(goString, "cty.", "", -1)

	return name, nil
}
